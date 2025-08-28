/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clusterinstance

import (
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// TestManagedFieldsHelper tests are integrated into the main ClusterInstanceSuite
// via the existing TestClusterInstance function in suite_test.go

var _ = Describe("ManagedFieldsHelper", func() {
	var (
		helper *ManagedFieldsHelper
		obj    *unstructured.Unstructured
		log    *zap.Logger
	)

	BeforeEach(func() {
		log = zap.NewNop() // Use no-op logger for tests
		helper = NewManagedFieldsHelper(log)
		
		obj = &unstructured.Unstructured{}
		obj.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "test.io",
			Version: "v1",
			Kind:    "TestResource",
		})
		obj.SetName("test-resource")
		obj.SetNamespace("test-namespace")
	})

	Describe("ValidateFieldManager", func() {
		It("should accept valid field manager names", func() {
			err := ValidateFieldManager("siteconfig-controller")
			Expect(err).ToNot(HaveOccurred())
			
			err = ValidateFieldManager("my-custom-controller")
			Expect(err).ToNot(HaveOccurred())
		})

		It("should reject empty field manager names", func() {
			err := ValidateFieldManager("")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("cannot be empty"))
		})

		It("should reject conflicting field manager names", func() {
			conflictingNames := []string{
				"kubectl",
				"kubectl-client-side-apply",
				"argocd-controller",
				"controller-manager",
			}

			for _, name := range conflictingNames {
				err := ValidateFieldManager(name)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("conflicts with well-known manager"))
			}
		})
	})

	Describe("GetFieldOwnership", func() {
		It("should handle objects without managedFields", func() {
			ownerships, err := helper.GetFieldOwnership(obj)
			Expect(err).ToNot(HaveOccurred())
			Expect(ownerships).To(BeEmpty())
		})

		It("should parse managedFields correctly", func() {
			// Create sample managedFields data
			fieldsV1Raw := `{
				"f:metadata": {
					"f:labels": {
						"f:test-label": {}
					},
					"f:annotations": {
						"f:test-annotation": {}
					}
				}
			}`

			managedFields := []metav1.ManagedFieldsEntry{
				{
					Manager:    "siteconfig-controller",
					Operation:  metav1.ManagedFieldsOperationApply,
					APIVersion: "v1",
					FieldsV1:   &metav1.FieldsV1{Raw: []byte(fieldsV1Raw)},
				},
			}

			obj.SetManagedFields(managedFields)

			ownerships, err := helper.GetFieldOwnership(obj)
			Expect(err).ToNot(HaveOccurred())
			Expect(ownerships).To(HaveLen(1))
			
			ownership := ownerships[0]
			Expect(ownership.Manager).To(Equal("siteconfig-controller"))
			Expect(ownership.Operation).To(Equal("Apply"))
			Expect(ownership.Fields).ToNot(BeEmpty())
		})
	})

	Describe("IsFieldManagedBy", func() {
		BeforeEach(func() {
			// Setup managedFields for testing
			fieldsV1Raw := `{
				"f:metadata": {
					"f:labels": {
						"f:managed-label": {}
					}
				}
			}`

			managedFields := []metav1.ManagedFieldsEntry{
				{
					Manager:    "siteconfig-controller",
					Operation:  metav1.ManagedFieldsOperationApply,
					APIVersion: "v1",
					FieldsV1:   &metav1.FieldsV1{Raw: []byte(fieldsV1Raw)},
				},
			}

			obj.SetManagedFields(managedFields)
		})

		It("should return true for fields managed by the specified manager", func() {
			// This is a simplified test - in practice, field path parsing would be more sophisticated
			result := helper.IsFieldManagedBy(obj, "f:metadata", "siteconfig-controller")
			Expect(result).To(BeTrue())
		})

		It("should return false for fields not managed by the specified manager", func() {
			result := helper.IsFieldManagedBy(obj, "f:metadata", "other-controller")
			Expect(result).To(BeFalse())
		})
	})

	Describe("LogFieldOwnership", func() {
		It("should not panic when logging field ownership", func() {
			// This test ensures the logging function doesn't panic
			Expect(func() {
				helper.LogFieldOwnership(obj)
			}).ToNot(Panic())
		})
	})
})
