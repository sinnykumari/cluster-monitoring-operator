// Copyright 2022 The Cluster Monitoring Operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package asciidocs

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/openshift/cluster-monitoring-operator/hack/docgen/model"
)

const (
	commonHeaders = `:_content-type: ASSEMBLY
include::_attributes/common-attributes.adoc[]
:context: configmap-reference-for-cluster-monitoring-operator
`
	firstParagraph = commonHeaders + `
[id="configmap-reference-for-cluster-monitoring-operator"]
= ConfigMap reference for Cluster Monitoring Operator

toc::[]

[id="cluster-monitoring-configuration-reference"]
== Cluster Monitoring configuration reference

Parts of Cluster Monitoring are configurable. Depending on which part of the stack users want to configure, they should edit the following:

- Configuration of OpenShift Container Platform monitoring components lies in a ConfigMap called ` + "`cluster-monitoring-config`" + ` in the ` + "`openshift-monitoring`" + ` namespace. Defined by link:#clustermonitoringconfiguration[ClusterMonitoringConfiguration].
- Configuration of components that monitor user-defined projects lies in a ConfigMap called ` + "`user-workload-monitoring-config`" + ` in the ` + "`openshift-user-workload-monitoring`" + ` namespace. Defined by link:#userworkloadconfiguration[UserWorkloadConfiguration].

The configuration file itself is always defined under the ` + "`config.yaml`" + ` key within the ConfigMap's data.

NOTE: Not all configuration parameters are exposed. Configuring Cluster Monitoring is optional. If the configuration does not exist or is empty or malformed, defaults are used.`
	pathToDocs    = "Documentation/openshiftdocs/"
	indexFile     = "index.adoc"
	modulesFolder = "modules/"
)

func toSectionLink(name string) string {
	name = strings.ToLower(name)
	name = strings.Replace(name, " ", "-", -1)
	return name
}

func PrintAPIDocs(args []string) {
	// TODO JoaoBraveCoding create dirs
	model.SetFormating("asciidoc")

	var indexContent string
	indexContent += firstParagraph

	paths := args[2:]
	typeSetUnion := make(model.TypeSet)
	typeSets := make([]model.TypeSet, 0, len(paths))
	for _, path := range paths {
		typeSet, err := model.Load(path)
		if err != nil {
			log.Fatal(err)
		}

		typeSets = append(typeSets, typeSet)
		for k, v := range typeSet {
			typeSetUnion[k] = v
		}
	}

	indexContent += "\n\n=== Table of Contents\n\n"

	for _, typeSet := range typeSets {
		for _, key := range typeSet.SortedKeys() {
			t := typeSet[key]
			if len(t.Fields) == 0 {
				continue
			}

			indexContent += fmt.Sprintf("* link:modules/%s.adoc[%s]\n", toSectionLink(t.Name), t.Name)
		}
	}

	writeToFile(pathToDocs+indexFile, indexContent)

	for _, typeSet := range typeSets {
		for _, key := range typeSet.SortedKeys() {
			t := typeSet[key]
			if len(t.Fields) == 0 {
				continue
			}

			// TODO JoaoBraveCoding create dirs
			var moduleContent string
			moduleContent += commonHeaders
			moduleContent += fmt.Sprintf("\n== %s\n\n=== Description\n\n%s\n\n", t.Name, t.Description())
			moduleContent += requiredSection(t) + "\n"

			backlinks := getBacklinks(t, typeSetUnion)
			if len(backlinks) > 0 {
				moduleContent += fmt.Sprintf("\nAppears in: %s\n\n", strings.Join(backlinks, ",\n"))
			}

			moduleContent += "[options=\"header\"]\n"
			moduleContent += "|===\n"
			moduleContent += "| Property | Type | Description \n"
			for _, f := range t.Fields {
				if strings.HasPrefix(fmt.Sprint(f.Description()), "OmitFromDoc") {
					continue
				}
				moduleContent += fmt.Sprint("|", f.Name(), "|", f.TypeLink(typeSetUnion), "|", f.Description(), "\n\n")
			}
			moduleContent += "|===\n"

			moduleContent += fmt.Sprintf("\nlink:%s[Back to TOC]\n", "../"+indexFile)

			writeToFile(pathToDocs+modulesFolder+toSectionLink(t.Name)+".adoc", moduleContent)
		}
	}
}

func requiredSection(t *model.StructType) string {
	var content string

	// Check if sctruct has any required fields
	hasRequiredFields := false
	for _, f := range t.Fields {
		if f.IsRequired() == true {
			hasRequiredFields = true
			break
		}
	}

	if hasRequiredFields {
		content += "=== Required\n"
		for _, f := range t.Fields {
			if f.IsRequired() == true {
				content += fmt.Sprint("* `", f.Name(), "`\n")
			}
		}
	}

	return content
}

func getBacklinks(t *model.StructType, typeSet model.TypeSet) []string {
	appearsIn := make(map[string]struct{})
	for _, v := range typeSet {
		if v.IsOnlyEmbedded() {
			continue
		}

		for _, f := range v.Fields {
			if f.TypeName() == t.Name {
				appearsIn[v.Name] = struct{}{}
			}
		}
	}

	var backlinks []string
	for item := range appearsIn {
		link := fmt.Sprintf("link:%s.adoc[%s]", toSectionLink(item), item)
		backlinks = append(backlinks, link)
	}
	sort.Strings(backlinks)

	return backlinks
}

func writeToFile(path, content string) {
	index, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	_, err = index.Write([]byte(content))
	if err != nil {
		log.Fatal(err)
	}
}
