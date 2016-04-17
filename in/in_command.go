package in

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/pivotal-cf-experimental/pivnet-resource/concourse"
	"github.com/pivotal-cf-experimental/pivnet-resource/downloader"
	"github.com/pivotal-cf-experimental/pivnet-resource/filter"
	"github.com/pivotal-cf-experimental/pivnet-resource/logger"
	"github.com/pivotal-cf-experimental/pivnet-resource/md5"
	"github.com/pivotal-cf-experimental/pivnet-resource/metadata"
	"github.com/pivotal-cf-experimental/pivnet-resource/pivnet"
	"github.com/pivotal-cf-experimental/pivnet-resource/versions"
)

type InCommand struct {
	logger       logger.Logger
	downloadDir  string
	pivnetClient pivnet.Client
	filter       filter.Filter
	downloader   downloader.Downloader
}

func NewInCommand(
	logger logger.Logger,
	downloadDir string,
	pivnetClient pivnet.Client,
	filter filter.Filter,
	downloader downloader.Downloader,
) *InCommand {
	return &InCommand{
		logger:       logger,
		downloadDir:  downloadDir,
		pivnetClient: pivnetClient,
		filter:       filter,
		downloader:   downloader,
	}
}

func (c *InCommand) Run(input concourse.InRequest) (concourse.InResponse, error) {
	c.logger.Debugf("Received input: %+v\n", input)

	productSlug := input.Source.ProductSlug

	c.logger.Debugf("Creating download directory: %s\n", c.downloadDir)
	err := os.MkdirAll(c.downloadDir, os.ModePerm)
	if err != nil {
		return concourse.InResponse{}, err
	}

	productVersion, etag, err := versions.SplitIntoVersionAndETag(input.Version.ProductVersion)
	if err != nil {
		c.logger.Debugf("Parsing of etag failed; continuing without it\n")
		productVersion = input.Version.ProductVersion
	}

	c.logger.Debugf(
		"Getting release: {product_slug: %s, product_version: %s, etag: %s}\n",
		productSlug,
		productVersion,
		etag,
	)

	release, err := c.pivnetClient.GetRelease(productSlug, productVersion)
	if err != nil {
		return concourse.InResponse{}, err
	}

	c.logger.Debugf(
		"Release: %+v\n",
		release,
	)

	c.logger.Debugf(
		"Accepting EULA: {product_slug: %s, release_id: %d}\n",
		productSlug,
		release.ID,
	)

	err = c.pivnetClient.AcceptEULA(productSlug, release.ID)
	if err != nil {
		return concourse.InResponse{}, err
	}

	c.logger.Debugf(
		"Getting product files: {release_id: %d}\n",
		release.ID,
	)

	productFiles, err := c.pivnetClient.GetProductFiles(release)
	if err != nil {
		return concourse.InResponse{}, err
	}

	c.logger.Debugf(
		"Getting download links: {product_files: %+v}\n",
		productFiles,
	)

	downloadLinksMD5 := map[string]string{}
	for _, p := range productFiles.ProductFiles {
		productFile, err := c.pivnetClient.GetProductFile(
			productSlug,
			release.ID,
			p.ID,
		)
		if err != nil {
			return concourse.InResponse{}, err
		}

		parts := strings.Split(productFile.AWSObjectKey, "/")
		fileName := parts[len(parts)-1]

		downloadLinksMD5[fileName] = productFile.MD5
	}

	downloadLinks := c.filter.DownloadLinks(productFiles)

	if len(input.Params.Globs) > 0 {
		c.logger.Debugf(
			"Filtering download links with globs: {globs: %+v}\n",
			input.Params.Globs,
		)

		var err error
		downloadLinks, err = c.filter.DownloadLinksByGlob(downloadLinks, input.Params.Globs)
		if err != nil {
			return concourse.InResponse{}, err
		}

		c.logger.Debugf(
			"Downloading files: {download_links: %+v, download_dir: %s}\n",
			downloadLinks,
			c.downloadDir,
		)

		files, err := c.downloader.Download(c.downloadDir, downloadLinks, input.Source.APIToken)
		if err != nil {
			log.Fatalf("Failed to Download Files: %s\n", err.Error())
		}

		for _, f := range files {
			downloadPath := filepath.Join(c.downloadDir, f)

			c.logger.Debugf(
				"Calcuating MD5 for downloaded file: %s\n",
				downloadPath,
			)
			md5, err := md5.NewFileContentsSummer(downloadPath).Sum()
			if err != nil {
				log.Fatalf("Failed to calculate MD5: %s\n", err.Error())
			}

			expectedMD5 := downloadLinksMD5[f]
			if md5 != expectedMD5 {
				log.Fatalf(
					"Failed MD5 comparison for file: %s. Expected %s, got %s\n",
					f,
					expectedMD5,
					md5,
				)
			}

			c.logger.Debugf(
				"MD5 for downloaded file: %s matched expected: %s\n",
				downloadPath,
				md5,
			)
		}
	}

	versionFilepath := filepath.Join(c.downloadDir, "version")

	versionWithETag, err := versions.CombineVersionAndETag(productVersion, etag)
	if err != nil {
		// Untested as it is too hard to force versions.CombineVersionAndETag
		// to return an error
		return concourse.InResponse{}, err
	}

	c.logger.Debugf(
		"Writing version to file: {version: %s, version_filepath: %s}\n",
		versionWithETag,
		versionFilepath,
	)

	err = ioutil.WriteFile(versionFilepath, []byte(versionWithETag), os.ModePerm)
	if err != nil {
		// Untested as it is too hard to force io.WriteFile to return an error
		return concourse.InResponse{}, err
	}

	mdata := metadata.Metadata{
		Release: &metadata.Release{
			Version:               release.Version,
			ReleaseType:           release.ReleaseType,
			ReleaseDate:           release.ReleaseDate,
			Description:           release.Description,
			ReleaseNotesURL:       release.ReleaseNotesURL,
			Availability:          release.Availability,
			Controlled:            release.Controlled,
			ECCN:                  release.ECCN,
			LicenseException:      release.LicenseException,
			EndOfSupportDate:      release.EndOfSupportDate,
			EndOfGuidanceDate:     release.EndOfGuidanceDate,
			EndOfAvailabilityDate: release.EndOfAvailabilityDate,
		},
	}

	for _, pf := range productFiles.ProductFiles {
		mdata.ProductFiles = append(mdata.ProductFiles, metadata.ProductFile{
			ID:           pf.ID,
			File:         pf.Name,
			Description:  pf.Description,
			AWSObjectKey: pf.AWSObjectKey,
			FileType:     pf.FileType,
			FileVersion:  pf.FileVersion,
			MD5:          pf.MD5,
		})
	}

	err = c.writeMetadataYAML(mdata)
	if err != nil {
		return concourse.InResponse{}, err
	}

	err = c.writeMetadataJSON(mdata)
	if err != nil {
		return concourse.InResponse{}, err
	}

	concourseMetadata := []concourse.Metadata{
		{Name: "version", Value: release.Version},
		{Name: "release_type", Value: release.ReleaseType},
		{Name: "release_date", Value: release.ReleaseDate},
		{Name: "description", Value: release.Description},
		{Name: "release_notes_url", Value: release.ReleaseNotesURL},
		{Name: "availability", Value: release.Availability},
		{Name: "controlled", Value: fmt.Sprintf("%t", release.Controlled)},
		{Name: "eccn", Value: release.ECCN},
		{Name: "license_exception", Value: release.LicenseException},
		{Name: "end_of_support_date", Value: release.EndOfSupportDate},
		{Name: "end_of_guidance_date", Value: release.EndOfGuidanceDate},
		{Name: "end_of_availability_date", Value: release.EndOfAvailabilityDate},
	}
	if release.EULA != nil {
		concourseMetadata = append(concourseMetadata, concourse.Metadata{Name: "eula_slug", Value: release.EULA.Slug})
	}

	if release.EULA != nil {
		concourseMetadata = append(concourseMetadata,
			concourse.Metadata{Name: "eula_slug", Value: release.EULA.Slug},
		)
	}

	out := concourse.InResponse{
		Version: concourse.Version{
			ProductVersion: versionWithETag,
		},
		Metadata: concourseMetadata,
	}

	return out, nil
}

func (c InCommand) writeMetadataJSON(mdata metadata.Metadata) error {
	jsonMetadataFilepath := filepath.Join(c.downloadDir, "metadata.json")
	c.logger.Debugf(
		"Writing metadata to json file: {metadata: %+v, metadata_filepath: %s}\n",
		mdata,
		jsonMetadataFilepath,
	)

	jsonMetadata, err := json.Marshal(mdata)
	if err != nil {
		// Untested as it is too hard to force json.Marshal to return an error
		return err
	}

	err = ioutil.WriteFile(jsonMetadataFilepath, jsonMetadata, os.ModePerm)
	if err != nil {
		// Untested as it is too hard to force io.WriteFile to return an error
		return err
	}

	return nil
}

func (c InCommand) writeMetadataYAML(mdata metadata.Metadata) error {
	yamlMetadataFilepath := filepath.Join(c.downloadDir, "metadata.yaml")
	c.logger.Debugf(
		"Writing metadata to yaml file: {metadata: %+v, metadata_filepath: %s}\n",
		mdata,
		yamlMetadataFilepath,
	)

	yamlMetadata, err := yaml.Marshal(mdata)
	if err != nil {
		// Untested as it is too hard to force yaml.Marshal to return an error
		return err
	}

	err = ioutil.WriteFile(yamlMetadataFilepath, yamlMetadata, os.ModePerm)
	if err != nil {
		// Untested as it is too hard to force io.WriteFile to return an error
		return err
	}

	return nil
}
