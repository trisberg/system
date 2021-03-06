/*
Copyright 2018 The Knative Authors

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

package resources

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	buildv1alpha1 "github.com/projectriff/system/pkg/apis/build/v1alpha1"
)

const (
	testFunctionName         = "test-function"
	testFunctionNamespace    = "test-function-namespace"
	testFunctionBuildName    = "test-function-function"
	testFunctionArtifact     = "test-function-artifact"
	testFunctionHandler      = "test-function-handler"
	testFunctionInvoker      = "test-function-invoker"
	testLabelKey             = "test-label-key"
	testLabelValue           = "test-label-value"
	testImage                = "test-image"
	testGitURL               = "https://example.com/repo.git"
	testGitRevision          = "master"
	testBuildCacheSize       = "8Gi"
	testBuildCacheName       = "build-cache-test-function-function"
	testApplicationName      = "test-application"
	testApplicationBuildName = "test-application-application"
	testApplicationNamespace = "test-application-namespace"
)

func TestFunctionBuild(t *testing.T) {
	f := &buildv1alpha1.Function{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testFunctionName,
			Namespace: testFunctionNamespace,
		},
	}
	f.Labels = map[string]string{testLabelKey: testLabelValue}
	f.Status.TargetImage = testImage
	f.Spec.Source = &buildv1alpha1.Source{
		Git: &buildv1alpha1.GitSource{
			URL:      testGitURL,
			Revision: testGitRevision,
		},
	}
	f.Status.BuildCacheName = testBuildCacheName
	f.Spec.Artifact = testFunctionArtifact
	f.Spec.Handler = testFunctionHandler
	f.Spec.Invoker = testFunctionInvoker

	b := MakeFunctionBuild(f)

	//if errs := b.Validate(context.Background()); errs != nil {
	//	t.Errorf("expected valid build got errors %+v", errs)
	//}

	if got, want := b.Name, testFunctionBuildName; got != want {
		t.Errorf("expected %q for build name got %q", want, got)
	}
	if got, want := b.Namespace, testFunctionNamespace; got != want {
		t.Errorf("expected %q for build namespace got %q", want, got)
	}

	if got, want := len(b.Labels), 2; got != want {
		t.Errorf("expected %d labels got %d", want, got)
	}
	if got, want := b.Labels[testLabelKey], testLabelValue; got != want {
		t.Errorf("expected %q labels got %q", want, got)
	}
	if got, want := b.Labels[buildv1alpha1.FunctionLabelKey], testFunctionName; got != want {
		t.Errorf("expected %q labels got %q", want, got)
	}

	if got, want := b.Spec.Template.Name, "riff-function"; got != want {
		t.Errorf("expected %q for template name got %q", want, got)
	}
	if got, want := len(b.Spec.Template.Arguments), 5; got != want {
		t.Errorf("expected %q template arg got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[0].Name, "IMAGE"; got != want {
		t.Errorf("expected %q for template image arg name got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[0].Value, testImage; got != want {
		t.Errorf("expected %q for template image arg value got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[1].Name, "FUNCTION_ARTIFACT"; got != want {
		t.Errorf("expected %q for function artifact arg name got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[1].Value, testFunctionArtifact; got != want {
		t.Errorf("expected %q for function artifact arg value got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[2].Name, "FUNCTION_HANDLER"; got != want {
		t.Errorf("expected %q for function handler arg name got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[2].Value, testFunctionHandler; got != want {
		t.Errorf("expected %q for function handler arg value got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[3].Name, "FUNCTION_LANGUAGE"; got != want {
		t.Errorf("expected %q for function language arg name got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[3].Value, testFunctionInvoker; got != want {
		t.Errorf("expected %q for function language arg value got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[4].Name, "CACHE"; got != want {
		t.Errorf("expected %q for template cache arg name got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[4].Value, "cache"; got != want {
		t.Errorf("expected %q for template cache arg value got %q", want, got)
	}

	if got, want := b.Spec.Source.Git.Url, testGitURL; got != want {
		t.Errorf("expected %q for git url got %q", want, got)
	}
	if got, want := b.Spec.Source.Git.Revision, testGitRevision; got != want {
		t.Errorf("expected %q for git revision got %q", want, got)
	}
}

func TestApplicationBuild(t *testing.T) {
	a := &buildv1alpha1.Application{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testApplicationName,
			Namespace: testApplicationNamespace,
		},
	}
	a.Labels = map[string]string{testLabelKey: testLabelValue}
	a.Status.TargetImage = testImage
	a.Spec.Source = &buildv1alpha1.Source{
		Git: &buildv1alpha1.GitSource{
			URL:      testGitURL,
			Revision: testGitRevision,
		},
	}
	a.Status.BuildCacheName = testBuildCacheName

	b := MakeApplicationBuild(a)

	//if errs := b.Validate(context.Background()); errs != nil {
	//	t.Errorf("expected valid build got errors %+v", errs)
	//}

	if got, want := b.Name, testApplicationBuildName; got != want {
		t.Errorf("expected %q for build name got %q", want, got)
	}
	if got, want := b.Namespace, testApplicationNamespace; got != want {
		t.Errorf("expected %q for build namespace got %q", want, got)
	}

	if got, want := len(b.Labels), 2; got != want {
		t.Errorf("expected %d labels got %d", want, got)
	}
	if got, want := b.Labels[testLabelKey], testLabelValue; got != want {
		t.Errorf("expected %q labels got %q", want, got)
	}
	if got, want := b.Labels[buildv1alpha1.ApplicationLabelKey], testApplicationName; got != want {
		t.Errorf("expected %q labels got %q", want, got)
	}

	if got, want := b.Spec.Template.Name, "riff-application"; got != want {
		t.Errorf("expected %q for template name got %q", want, got)
	}
	if got, want := len(b.Spec.Template.Arguments), 2; got != want {
		t.Errorf("expected %q template arg got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[0].Name, "IMAGE"; got != want {
		t.Errorf("expected %q for template image arg name got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[0].Value, testImage; got != want {
		t.Errorf("expected %q for template image arg value got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[1].Name, "CACHE"; got != want {
		t.Errorf("expected %q for template cache arg name got %q", want, got)
	}
	if got, want := b.Spec.Template.Arguments[1].Value, "cache"; got != want {
		t.Errorf("expected %q for template cache arg value got %q", want, got)
	}

	if got, want := b.Spec.Source.Git.Url, testGitURL; got != want {
		t.Errorf("expected %q for git url got %q", want, got)
	}
	if got, want := b.Spec.Source.Git.Revision, testGitRevision; got != want {
		t.Errorf("expected %q for git revision got %q", want, got)
	}
}
