// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package codespace

import (
	"context"
	"net/http"
	"sync"

	codespacesAPI "github.com/cli/cli/v2/internal/codespaces/api"
)

// apiClientMock is a mock implementation of apiClient.
//
//	func TestSomethingThatUsesapiClient(t *testing.T) {
//
//		// make and configure a mocked apiClient
//		mockedapiClient := &apiClientMock{
//			CreateCodespaceFunc: func(ctx context.Context, params *codespacesAPI.CreateCodespaceParams) (*codespacesAPI.Codespace, error) {
//				panic("mock out the CreateCodespace method")
//			},
//			DeleteCodespaceFunc: func(ctx context.Context, name string, orgName string, userName string) error {
//				panic("mock out the DeleteCodespace method")
//			},
//			EditCodespaceFunc: func(ctx context.Context, codespaceName string, params *codespacesAPI.EditCodespaceParams) (*codespacesAPI.Codespace, error) {
//				panic("mock out the EditCodespace method")
//			},
//			GetCodespaceFunc: func(ctx context.Context, name string, includeConnection bool) (*codespacesAPI.Codespace, error) {
//				panic("mock out the GetCodespace method")
//			},
//			GetCodespaceBillableOwnerFunc: func(ctx context.Context, nwo string) (*codespacesAPI.User, error) {
//				panic("mock out the GetCodespaceBillableOwner method")
//			},
//			GetCodespaceRepoSuggestionsFunc: func(ctx context.Context, partialSearch string, params codespacesAPI.RepoSearchParameters) ([]string, error) {
//				panic("mock out the GetCodespaceRepoSuggestions method")
//			},
//			GetCodespaceRepositoryContentsFunc: func(ctx context.Context, codespace *codespacesAPI.Codespace, path string) ([]byte, error) {
//				panic("mock out the GetCodespaceRepositoryContents method")
//			},
//			GetCodespacesMachinesFunc: func(ctx context.Context, repoID int, branch string, location string, devcontainerPath string) ([]*codespacesAPI.Machine, error) {
//				panic("mock out the GetCodespacesMachines method")
//			},
//			GetCodespacesPermissionsCheckFunc: func(ctx context.Context, repoID int, branch string, devcontainerPath string) (bool, error) {
//				panic("mock out the GetCodespacesPermissionsCheck method")
//			},
//			GetOrgMemberCodespaceFunc: func(ctx context.Context, orgName string, userName string, codespaceName string) (*codespacesAPI.Codespace, error) {
//				panic("mock out the GetOrgMemberCodespace method")
//			},
//			GetRepositoryFunc: func(ctx context.Context, nwo string) (*codespacesAPI.Repository, error) {
//				panic("mock out the GetRepository method")
//			},
//			GetUserFunc: func(ctx context.Context) (*codespacesAPI.User, error) {
//				panic("mock out the GetUser method")
//			},
//			HTTPClientFunc: func() (*http.Client, error) {
//				panic("mock out the HTTPClient method")
//			},
//			ListCodespacesFunc: func(ctx context.Context, opts codespacesAPI.ListCodespacesOptions) ([]*codespacesAPI.Codespace, error) {
//				panic("mock out the ListCodespaces method")
//			},
//			ListDevContainersFunc: func(ctx context.Context, repoID int, branch string, limit int) ([]codespacesAPI.DevContainerEntry, error) {
//				panic("mock out the ListDevContainers method")
//			},
//			ServerURLFunc: func() string {
//				panic("mock out the ServerURL method")
//			},
//			StartCodespaceFunc: func(ctx context.Context, name string) error {
//				panic("mock out the StartCodespace method")
//			},
//			StopCodespaceFunc: func(ctx context.Context, name string, orgName string, userName string) error {
//				panic("mock out the StopCodespace method")
//			},
//		}
//
//		// use mockedapiClient in code that requires apiClient
//		// and then make assertions.
//
//	}
type apiClientMock struct {
	// CreateCodespaceFunc mocks the CreateCodespace method.
	CreateCodespaceFunc func(ctx context.Context, params *codespacesAPI.CreateCodespaceParams) (*codespacesAPI.Codespace, error)

	// DeleteCodespaceFunc mocks the DeleteCodespace method.
	DeleteCodespaceFunc func(ctx context.Context, name string, orgName string, userName string) error

	// EditCodespaceFunc mocks the EditCodespace method.
	EditCodespaceFunc func(ctx context.Context, codespaceName string, params *codespacesAPI.EditCodespaceParams) (*codespacesAPI.Codespace, error)

	// GetCodespaceFunc mocks the GetCodespace method.
	GetCodespaceFunc func(ctx context.Context, name string, includeConnection bool) (*codespacesAPI.Codespace, error)

	// GetCodespaceBillableOwnerFunc mocks the GetCodespaceBillableOwner method.
	GetCodespaceBillableOwnerFunc func(ctx context.Context, nwo string) (*codespacesAPI.User, error)

	// GetCodespaceRepoSuggestionsFunc mocks the GetCodespaceRepoSuggestions method.
	GetCodespaceRepoSuggestionsFunc func(ctx context.Context, partialSearch string, params codespacesAPI.RepoSearchParameters) ([]string, error)

	// GetCodespaceRepositoryContentsFunc mocks the GetCodespaceRepositoryContents method.
	GetCodespaceRepositoryContentsFunc func(ctx context.Context, codespace *codespacesAPI.Codespace, path string) ([]byte, error)

	// GetCodespacesMachinesFunc mocks the GetCodespacesMachines method.
	GetCodespacesMachinesFunc func(ctx context.Context, repoID int, branch string, location string, devcontainerPath string) ([]*codespacesAPI.Machine, error)

	// GetCodespacesPermissionsCheckFunc mocks the GetCodespacesPermissionsCheck method.
	GetCodespacesPermissionsCheckFunc func(ctx context.Context, repoID int, branch string, devcontainerPath string) (bool, error)

	// GetOrgMemberCodespaceFunc mocks the GetOrgMemberCodespace method.
	GetOrgMemberCodespaceFunc func(ctx context.Context, orgName string, userName string, codespaceName string) (*codespacesAPI.Codespace, error)

	// GetRepositoryFunc mocks the GetRepository method.
	GetRepositoryFunc func(ctx context.Context, nwo string) (*codespacesAPI.Repository, error)

	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(ctx context.Context) (*codespacesAPI.User, error)

	// HTTPClientFunc mocks the HTTPClient method.
	HTTPClientFunc func() (*http.Client, error)

	// ListCodespacesFunc mocks the ListCodespaces method.
	ListCodespacesFunc func(ctx context.Context, opts codespacesAPI.ListCodespacesOptions) ([]*codespacesAPI.Codespace, error)

	// ListDevContainersFunc mocks the ListDevContainers method.
	ListDevContainersFunc func(ctx context.Context, repoID int, branch string, limit int) ([]codespacesAPI.DevContainerEntry, error)

	// ServerURLFunc mocks the ServerURL method.
	ServerURLFunc func() string

	// StartCodespaceFunc mocks the StartCodespace method.
	StartCodespaceFunc func(ctx context.Context, name string) error

	// StopCodespaceFunc mocks the StopCodespace method.
	StopCodespaceFunc func(ctx context.Context, name string, orgName string, userName string) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateCodespace holds details about calls to the CreateCodespace method.
		CreateCodespace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *codespacesAPI.CreateCodespaceParams
		}
		// DeleteCodespace holds details about calls to the DeleteCodespace method.
		DeleteCodespace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// OrgName is the orgName argument value.
			OrgName string
			// UserName is the userName argument value.
			UserName string
		}
		// EditCodespace holds details about calls to the EditCodespace method.
		EditCodespace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CodespaceName is the codespaceName argument value.
			CodespaceName string
			// Params is the params argument value.
			Params *codespacesAPI.EditCodespaceParams
		}
		// GetCodespace holds details about calls to the GetCodespace method.
		GetCodespace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// IncludeConnection is the includeConnection argument value.
			IncludeConnection bool
		}
		// GetCodespaceBillableOwner holds details about calls to the GetCodespaceBillableOwner method.
		GetCodespaceBillableOwner []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Nwo is the nwo argument value.
			Nwo string
		}
		// GetCodespaceRepoSuggestions holds details about calls to the GetCodespaceRepoSuggestions method.
		GetCodespaceRepoSuggestions []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PartialSearch is the partialSearch argument value.
			PartialSearch string
			// Params is the params argument value.
			Params codespacesAPI.RepoSearchParameters
		}
		// GetCodespaceRepositoryContents holds details about calls to the GetCodespaceRepositoryContents method.
		GetCodespaceRepositoryContents []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Codespace is the codespace argument value.
			Codespace *codespacesAPI.Codespace
			// Path is the path argument value.
			Path string
		}
		// GetCodespacesMachines holds details about calls to the GetCodespacesMachines method.
		GetCodespacesMachines []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// RepoID is the repoID argument value.
			RepoID int
			// Branch is the branch argument value.
			Branch string
			// Location is the location argument value.
			Location string
			// DevcontainerPath is the devcontainerPath argument value.
			DevcontainerPath string
		}
		// GetCodespacesPermissionsCheck holds details about calls to the GetCodespacesPermissionsCheck method.
		GetCodespacesPermissionsCheck []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// RepoID is the repoID argument value.
			RepoID int
			// Branch is the branch argument value.
			Branch string
			// DevcontainerPath is the devcontainerPath argument value.
			DevcontainerPath string
		}
		// GetOrgMemberCodespace holds details about calls to the GetOrgMemberCodespace method.
		GetOrgMemberCodespace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OrgName is the orgName argument value.
			OrgName string
			// UserName is the userName argument value.
			UserName string
			// CodespaceName is the codespaceName argument value.
			CodespaceName string
		}
		// GetRepository holds details about calls to the GetRepository method.
		GetRepository []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Nwo is the nwo argument value.
			Nwo string
		}
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// HTTPClient holds details about calls to the HTTPClient method.
		HTTPClient []struct {
		}
		// ListCodespaces holds details about calls to the ListCodespaces method.
		ListCodespaces []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts codespacesAPI.ListCodespacesOptions
		}
		// ListDevContainers holds details about calls to the ListDevContainers method.
		ListDevContainers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// RepoID is the repoID argument value.
			RepoID int
			// Branch is the branch argument value.
			Branch string
			// Limit is the limit argument value.
			Limit int
		}
		// ServerURL holds details about calls to the ServerURL method.
		ServerURL []struct {
		}
		// StartCodespace holds details about calls to the StartCodespace method.
		StartCodespace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// StopCodespace holds details about calls to the StopCodespace method.
		StopCodespace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// OrgName is the orgName argument value.
			OrgName string
			// UserName is the userName argument value.
			UserName string
		}
	}
	lockCreateCodespace                sync.RWMutex
	lockDeleteCodespace                sync.RWMutex
	lockEditCodespace                  sync.RWMutex
	lockGetCodespace                   sync.RWMutex
	lockGetCodespaceBillableOwner      sync.RWMutex
	lockGetCodespaceRepoSuggestions    sync.RWMutex
	lockGetCodespaceRepositoryContents sync.RWMutex
	lockGetCodespacesMachines          sync.RWMutex
	lockGetCodespacesPermissionsCheck  sync.RWMutex
	lockGetOrgMemberCodespace          sync.RWMutex
	lockGetRepository                  sync.RWMutex
	lockGetUser                        sync.RWMutex
	lockHTTPClient                     sync.RWMutex
	lockListCodespaces                 sync.RWMutex
	lockListDevContainers              sync.RWMutex
	lockServerURL                      sync.RWMutex
	lockStartCodespace                 sync.RWMutex
	lockStopCodespace                  sync.RWMutex
}

// CreateCodespace calls CreateCodespaceFunc.
func (mock *apiClientMock) CreateCodespace(ctx context.Context, params *codespacesAPI.CreateCodespaceParams) (*codespacesAPI.Codespace, error) {
	if mock.CreateCodespaceFunc == nil {
		panic("apiClientMock.CreateCodespaceFunc: method is nil but apiClient.CreateCodespace was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *codespacesAPI.CreateCodespaceParams
	}{
		Ctx:    ctx,
		Params: params,
	}
	mock.lockCreateCodespace.Lock()
	mock.calls.CreateCodespace = append(mock.calls.CreateCodespace, callInfo)
	mock.lockCreateCodespace.Unlock()
	return mock.CreateCodespaceFunc(ctx, params)
}

// CreateCodespaceCalls gets all the calls that were made to CreateCodespace.
// Check the length with:
//
//	len(mockedapiClient.CreateCodespaceCalls())
func (mock *apiClientMock) CreateCodespaceCalls() []struct {
	Ctx    context.Context
	Params *codespacesAPI.CreateCodespaceParams
} {
	var calls []struct {
		Ctx    context.Context
		Params *codespacesAPI.CreateCodespaceParams
	}
	mock.lockCreateCodespace.RLock()
	calls = mock.calls.CreateCodespace
	mock.lockCreateCodespace.RUnlock()
	return calls
}

// DeleteCodespace calls DeleteCodespaceFunc.
func (mock *apiClientMock) DeleteCodespace(ctx context.Context, name string, orgName string, userName string) error {
	if mock.DeleteCodespaceFunc == nil {
		panic("apiClientMock.DeleteCodespaceFunc: method is nil but apiClient.DeleteCodespace was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Name     string
		OrgName  string
		UserName string
	}{
		Ctx:      ctx,
		Name:     name,
		OrgName:  orgName,
		UserName: userName,
	}
	mock.lockDeleteCodespace.Lock()
	mock.calls.DeleteCodespace = append(mock.calls.DeleteCodespace, callInfo)
	mock.lockDeleteCodespace.Unlock()
	return mock.DeleteCodespaceFunc(ctx, name, orgName, userName)
}

// DeleteCodespaceCalls gets all the calls that were made to DeleteCodespace.
// Check the length with:
//
//	len(mockedapiClient.DeleteCodespaceCalls())
func (mock *apiClientMock) DeleteCodespaceCalls() []struct {
	Ctx      context.Context
	Name     string
	OrgName  string
	UserName string
} {
	var calls []struct {
		Ctx      context.Context
		Name     string
		OrgName  string
		UserName string
	}
	mock.lockDeleteCodespace.RLock()
	calls = mock.calls.DeleteCodespace
	mock.lockDeleteCodespace.RUnlock()
	return calls
}

// EditCodespace calls EditCodespaceFunc.
func (mock *apiClientMock) EditCodespace(ctx context.Context, codespaceName string, params *codespacesAPI.EditCodespaceParams) (*codespacesAPI.Codespace, error) {
	if mock.EditCodespaceFunc == nil {
		panic("apiClientMock.EditCodespaceFunc: method is nil but apiClient.EditCodespace was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		CodespaceName string
		Params        *codespacesAPI.EditCodespaceParams
	}{
		Ctx:           ctx,
		CodespaceName: codespaceName,
		Params:        params,
	}
	mock.lockEditCodespace.Lock()
	mock.calls.EditCodespace = append(mock.calls.EditCodespace, callInfo)
	mock.lockEditCodespace.Unlock()
	return mock.EditCodespaceFunc(ctx, codespaceName, params)
}

// EditCodespaceCalls gets all the calls that were made to EditCodespace.
// Check the length with:
//
//	len(mockedapiClient.EditCodespaceCalls())
func (mock *apiClientMock) EditCodespaceCalls() []struct {
	Ctx           context.Context
	CodespaceName string
	Params        *codespacesAPI.EditCodespaceParams
} {
	var calls []struct {
		Ctx           context.Context
		CodespaceName string
		Params        *codespacesAPI.EditCodespaceParams
	}
	mock.lockEditCodespace.RLock()
	calls = mock.calls.EditCodespace
	mock.lockEditCodespace.RUnlock()
	return calls
}

// GetCodespace calls GetCodespaceFunc.
func (mock *apiClientMock) GetCodespace(ctx context.Context, name string, includeConnection bool) (*codespacesAPI.Codespace, error) {
	if mock.GetCodespaceFunc == nil {
		panic("apiClientMock.GetCodespaceFunc: method is nil but apiClient.GetCodespace was just called")
	}
	callInfo := struct {
		Ctx               context.Context
		Name              string
		IncludeConnection bool
	}{
		Ctx:               ctx,
		Name:              name,
		IncludeConnection: includeConnection,
	}
	mock.lockGetCodespace.Lock()
	mock.calls.GetCodespace = append(mock.calls.GetCodespace, callInfo)
	mock.lockGetCodespace.Unlock()
	return mock.GetCodespaceFunc(ctx, name, includeConnection)
}

// GetCodespaceCalls gets all the calls that were made to GetCodespace.
// Check the length with:
//
//	len(mockedapiClient.GetCodespaceCalls())
func (mock *apiClientMock) GetCodespaceCalls() []struct {
	Ctx               context.Context
	Name              string
	IncludeConnection bool
} {
	var calls []struct {
		Ctx               context.Context
		Name              string
		IncludeConnection bool
	}
	mock.lockGetCodespace.RLock()
	calls = mock.calls.GetCodespace
	mock.lockGetCodespace.RUnlock()
	return calls
}

// GetCodespaceBillableOwner calls GetCodespaceBillableOwnerFunc.
func (mock *apiClientMock) GetCodespaceBillableOwner(ctx context.Context, nwo string) (*codespacesAPI.User, error) {
	if mock.GetCodespaceBillableOwnerFunc == nil {
		panic("apiClientMock.GetCodespaceBillableOwnerFunc: method is nil but apiClient.GetCodespaceBillableOwner was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Nwo string
	}{
		Ctx: ctx,
		Nwo: nwo,
	}
	mock.lockGetCodespaceBillableOwner.Lock()
	mock.calls.GetCodespaceBillableOwner = append(mock.calls.GetCodespaceBillableOwner, callInfo)
	mock.lockGetCodespaceBillableOwner.Unlock()
	return mock.GetCodespaceBillableOwnerFunc(ctx, nwo)
}

// GetCodespaceBillableOwnerCalls gets all the calls that were made to GetCodespaceBillableOwner.
// Check the length with:
//
//	len(mockedapiClient.GetCodespaceBillableOwnerCalls())
func (mock *apiClientMock) GetCodespaceBillableOwnerCalls() []struct {
	Ctx context.Context
	Nwo string
} {
	var calls []struct {
		Ctx context.Context
		Nwo string
	}
	mock.lockGetCodespaceBillableOwner.RLock()
	calls = mock.calls.GetCodespaceBillableOwner
	mock.lockGetCodespaceBillableOwner.RUnlock()
	return calls
}

// GetCodespaceRepoSuggestions calls GetCodespaceRepoSuggestionsFunc.
func (mock *apiClientMock) GetCodespaceRepoSuggestions(ctx context.Context, partialSearch string, params codespacesAPI.RepoSearchParameters) ([]string, error) {
	if mock.GetCodespaceRepoSuggestionsFunc == nil {
		panic("apiClientMock.GetCodespaceRepoSuggestionsFunc: method is nil but apiClient.GetCodespaceRepoSuggestions was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		PartialSearch string
		Params        codespacesAPI.RepoSearchParameters
	}{
		Ctx:           ctx,
		PartialSearch: partialSearch,
		Params:        params,
	}
	mock.lockGetCodespaceRepoSuggestions.Lock()
	mock.calls.GetCodespaceRepoSuggestions = append(mock.calls.GetCodespaceRepoSuggestions, callInfo)
	mock.lockGetCodespaceRepoSuggestions.Unlock()
	return mock.GetCodespaceRepoSuggestionsFunc(ctx, partialSearch, params)
}

// GetCodespaceRepoSuggestionsCalls gets all the calls that were made to GetCodespaceRepoSuggestions.
// Check the length with:
//
//	len(mockedapiClient.GetCodespaceRepoSuggestionsCalls())
func (mock *apiClientMock) GetCodespaceRepoSuggestionsCalls() []struct {
	Ctx           context.Context
	PartialSearch string
	Params        codespacesAPI.RepoSearchParameters
} {
	var calls []struct {
		Ctx           context.Context
		PartialSearch string
		Params        codespacesAPI.RepoSearchParameters
	}
	mock.lockGetCodespaceRepoSuggestions.RLock()
	calls = mock.calls.GetCodespaceRepoSuggestions
	mock.lockGetCodespaceRepoSuggestions.RUnlock()
	return calls
}

// GetCodespaceRepositoryContents calls GetCodespaceRepositoryContentsFunc.
func (mock *apiClientMock) GetCodespaceRepositoryContents(ctx context.Context, codespace *codespacesAPI.Codespace, path string) ([]byte, error) {
	if mock.GetCodespaceRepositoryContentsFunc == nil {
		panic("apiClientMock.GetCodespaceRepositoryContentsFunc: method is nil but apiClient.GetCodespaceRepositoryContents was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Codespace *codespacesAPI.Codespace
		Path      string
	}{
		Ctx:       ctx,
		Codespace: codespace,
		Path:      path,
	}
	mock.lockGetCodespaceRepositoryContents.Lock()
	mock.calls.GetCodespaceRepositoryContents = append(mock.calls.GetCodespaceRepositoryContents, callInfo)
	mock.lockGetCodespaceRepositoryContents.Unlock()
	return mock.GetCodespaceRepositoryContentsFunc(ctx, codespace, path)
}

// GetCodespaceRepositoryContentsCalls gets all the calls that were made to GetCodespaceRepositoryContents.
// Check the length with:
//
//	len(mockedapiClient.GetCodespaceRepositoryContentsCalls())
func (mock *apiClientMock) GetCodespaceRepositoryContentsCalls() []struct {
	Ctx       context.Context
	Codespace *codespacesAPI.Codespace
	Path      string
} {
	var calls []struct {
		Ctx       context.Context
		Codespace *codespacesAPI.Codespace
		Path      string
	}
	mock.lockGetCodespaceRepositoryContents.RLock()
	calls = mock.calls.GetCodespaceRepositoryContents
	mock.lockGetCodespaceRepositoryContents.RUnlock()
	return calls
}

// GetCodespacesMachines calls GetCodespacesMachinesFunc.
func (mock *apiClientMock) GetCodespacesMachines(ctx context.Context, repoID int, branch string, location string, devcontainerPath string) ([]*codespacesAPI.Machine, error) {
	if mock.GetCodespacesMachinesFunc == nil {
		panic("apiClientMock.GetCodespacesMachinesFunc: method is nil but apiClient.GetCodespacesMachines was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		RepoID           int
		Branch           string
		Location         string
		DevcontainerPath string
	}{
		Ctx:              ctx,
		RepoID:           repoID,
		Branch:           branch,
		Location:         location,
		DevcontainerPath: devcontainerPath,
	}
	mock.lockGetCodespacesMachines.Lock()
	mock.calls.GetCodespacesMachines = append(mock.calls.GetCodespacesMachines, callInfo)
	mock.lockGetCodespacesMachines.Unlock()
	return mock.GetCodespacesMachinesFunc(ctx, repoID, branch, location, devcontainerPath)
}

// GetCodespacesMachinesCalls gets all the calls that were made to GetCodespacesMachines.
// Check the length with:
//
//	len(mockedapiClient.GetCodespacesMachinesCalls())
func (mock *apiClientMock) GetCodespacesMachinesCalls() []struct {
	Ctx              context.Context
	RepoID           int
	Branch           string
	Location         string
	DevcontainerPath string
} {
	var calls []struct {
		Ctx              context.Context
		RepoID           int
		Branch           string
		Location         string
		DevcontainerPath string
	}
	mock.lockGetCodespacesMachines.RLock()
	calls = mock.calls.GetCodespacesMachines
	mock.lockGetCodespacesMachines.RUnlock()
	return calls
}

// GetCodespacesPermissionsCheck calls GetCodespacesPermissionsCheckFunc.
func (mock *apiClientMock) GetCodespacesPermissionsCheck(ctx context.Context, repoID int, branch string, devcontainerPath string) (bool, error) {
	if mock.GetCodespacesPermissionsCheckFunc == nil {
		panic("apiClientMock.GetCodespacesPermissionsCheckFunc: method is nil but apiClient.GetCodespacesPermissionsCheck was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		RepoID           int
		Branch           string
		DevcontainerPath string
	}{
		Ctx:              ctx,
		RepoID:           repoID,
		Branch:           branch,
		DevcontainerPath: devcontainerPath,
	}
	mock.lockGetCodespacesPermissionsCheck.Lock()
	mock.calls.GetCodespacesPermissionsCheck = append(mock.calls.GetCodespacesPermissionsCheck, callInfo)
	mock.lockGetCodespacesPermissionsCheck.Unlock()
	return mock.GetCodespacesPermissionsCheckFunc(ctx, repoID, branch, devcontainerPath)
}

// GetCodespacesPermissionsCheckCalls gets all the calls that were made to GetCodespacesPermissionsCheck.
// Check the length with:
//
//	len(mockedapiClient.GetCodespacesPermissionsCheckCalls())
func (mock *apiClientMock) GetCodespacesPermissionsCheckCalls() []struct {
	Ctx              context.Context
	RepoID           int
	Branch           string
	DevcontainerPath string
} {
	var calls []struct {
		Ctx              context.Context
		RepoID           int
		Branch           string
		DevcontainerPath string
	}
	mock.lockGetCodespacesPermissionsCheck.RLock()
	calls = mock.calls.GetCodespacesPermissionsCheck
	mock.lockGetCodespacesPermissionsCheck.RUnlock()
	return calls
}

// GetOrgMemberCodespace calls GetOrgMemberCodespaceFunc.
func (mock *apiClientMock) GetOrgMemberCodespace(ctx context.Context, orgName string, userName string, codespaceName string) (*codespacesAPI.Codespace, error) {
	if mock.GetOrgMemberCodespaceFunc == nil {
		panic("apiClientMock.GetOrgMemberCodespaceFunc: method is nil but apiClient.GetOrgMemberCodespace was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		OrgName       string
		UserName      string
		CodespaceName string
	}{
		Ctx:           ctx,
		OrgName:       orgName,
		UserName:      userName,
		CodespaceName: codespaceName,
	}
	mock.lockGetOrgMemberCodespace.Lock()
	mock.calls.GetOrgMemberCodespace = append(mock.calls.GetOrgMemberCodespace, callInfo)
	mock.lockGetOrgMemberCodespace.Unlock()
	return mock.GetOrgMemberCodespaceFunc(ctx, orgName, userName, codespaceName)
}

// GetOrgMemberCodespaceCalls gets all the calls that were made to GetOrgMemberCodespace.
// Check the length with:
//
//	len(mockedapiClient.GetOrgMemberCodespaceCalls())
func (mock *apiClientMock) GetOrgMemberCodespaceCalls() []struct {
	Ctx           context.Context
	OrgName       string
	UserName      string
	CodespaceName string
} {
	var calls []struct {
		Ctx           context.Context
		OrgName       string
		UserName      string
		CodespaceName string
	}
	mock.lockGetOrgMemberCodespace.RLock()
	calls = mock.calls.GetOrgMemberCodespace
	mock.lockGetOrgMemberCodespace.RUnlock()
	return calls
}

// GetRepository calls GetRepositoryFunc.
func (mock *apiClientMock) GetRepository(ctx context.Context, nwo string) (*codespacesAPI.Repository, error) {
	if mock.GetRepositoryFunc == nil {
		panic("apiClientMock.GetRepositoryFunc: method is nil but apiClient.GetRepository was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Nwo string
	}{
		Ctx: ctx,
		Nwo: nwo,
	}
	mock.lockGetRepository.Lock()
	mock.calls.GetRepository = append(mock.calls.GetRepository, callInfo)
	mock.lockGetRepository.Unlock()
	return mock.GetRepositoryFunc(ctx, nwo)
}

// GetRepositoryCalls gets all the calls that were made to GetRepository.
// Check the length with:
//
//	len(mockedapiClient.GetRepositoryCalls())
func (mock *apiClientMock) GetRepositoryCalls() []struct {
	Ctx context.Context
	Nwo string
} {
	var calls []struct {
		Ctx context.Context
		Nwo string
	}
	mock.lockGetRepository.RLock()
	calls = mock.calls.GetRepository
	mock.lockGetRepository.RUnlock()
	return calls
}

// GetUser calls GetUserFunc.
func (mock *apiClientMock) GetUser(ctx context.Context) (*codespacesAPI.User, error) {
	if mock.GetUserFunc == nil {
		panic("apiClientMock.GetUserFunc: method is nil but apiClient.GetUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	mock.lockGetUser.Unlock()
	return mock.GetUserFunc(ctx)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//
//	len(mockedapiClient.GetUserCalls())
func (mock *apiClientMock) GetUserCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetUser.RLock()
	calls = mock.calls.GetUser
	mock.lockGetUser.RUnlock()
	return calls
}

// HTTPClient calls HTTPClientFunc.
func (mock *apiClientMock) HTTPClient() (*http.Client, error) {
	if mock.HTTPClientFunc == nil {
		panic("apiClientMock.HTTPClientFunc: method is nil but apiClient.HTTPClient was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHTTPClient.Lock()
	mock.calls.HTTPClient = append(mock.calls.HTTPClient, callInfo)
	mock.lockHTTPClient.Unlock()
	return mock.HTTPClientFunc()
}

// HTTPClientCalls gets all the calls that were made to HTTPClient.
// Check the length with:
//
//	len(mockedapiClient.HTTPClientCalls())
func (mock *apiClientMock) HTTPClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHTTPClient.RLock()
	calls = mock.calls.HTTPClient
	mock.lockHTTPClient.RUnlock()
	return calls
}

// ListCodespaces calls ListCodespacesFunc.
func (mock *apiClientMock) ListCodespaces(ctx context.Context, opts codespacesAPI.ListCodespacesOptions) ([]*codespacesAPI.Codespace, error) {
	if mock.ListCodespacesFunc == nil {
		panic("apiClientMock.ListCodespacesFunc: method is nil but apiClient.ListCodespaces was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts codespacesAPI.ListCodespacesOptions
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockListCodespaces.Lock()
	mock.calls.ListCodespaces = append(mock.calls.ListCodespaces, callInfo)
	mock.lockListCodespaces.Unlock()
	return mock.ListCodespacesFunc(ctx, opts)
}

// ListCodespacesCalls gets all the calls that were made to ListCodespaces.
// Check the length with:
//
//	len(mockedapiClient.ListCodespacesCalls())
func (mock *apiClientMock) ListCodespacesCalls() []struct {
	Ctx  context.Context
	Opts codespacesAPI.ListCodespacesOptions
} {
	var calls []struct {
		Ctx  context.Context
		Opts codespacesAPI.ListCodespacesOptions
	}
	mock.lockListCodespaces.RLock()
	calls = mock.calls.ListCodespaces
	mock.lockListCodespaces.RUnlock()
	return calls
}

// ListDevContainers calls ListDevContainersFunc.
func (mock *apiClientMock) ListDevContainers(ctx context.Context, repoID int, branch string, limit int) ([]codespacesAPI.DevContainerEntry, error) {
	if mock.ListDevContainersFunc == nil {
		panic("apiClientMock.ListDevContainersFunc: method is nil but apiClient.ListDevContainers was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		RepoID int
		Branch string
		Limit  int
	}{
		Ctx:    ctx,
		RepoID: repoID,
		Branch: branch,
		Limit:  limit,
	}
	mock.lockListDevContainers.Lock()
	mock.calls.ListDevContainers = append(mock.calls.ListDevContainers, callInfo)
	mock.lockListDevContainers.Unlock()
	return mock.ListDevContainersFunc(ctx, repoID, branch, limit)
}

// ListDevContainersCalls gets all the calls that were made to ListDevContainers.
// Check the length with:
//
//	len(mockedapiClient.ListDevContainersCalls())
func (mock *apiClientMock) ListDevContainersCalls() []struct {
	Ctx    context.Context
	RepoID int
	Branch string
	Limit  int
} {
	var calls []struct {
		Ctx    context.Context
		RepoID int
		Branch string
		Limit  int
	}
	mock.lockListDevContainers.RLock()
	calls = mock.calls.ListDevContainers
	mock.lockListDevContainers.RUnlock()
	return calls
}

// ServerURL calls ServerURLFunc.
func (mock *apiClientMock) ServerURL() string {
	if mock.ServerURLFunc == nil {
		panic("apiClientMock.ServerURLFunc: method is nil but apiClient.ServerURL was just called")
	}
	callInfo := struct {
	}{}
	mock.lockServerURL.Lock()
	mock.calls.ServerURL = append(mock.calls.ServerURL, callInfo)
	mock.lockServerURL.Unlock()
	return mock.ServerURLFunc()
}

// ServerURLCalls gets all the calls that were made to ServerURL.
// Check the length with:
//
//	len(mockedapiClient.ServerURLCalls())
func (mock *apiClientMock) ServerURLCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockServerURL.RLock()
	calls = mock.calls.ServerURL
	mock.lockServerURL.RUnlock()
	return calls
}

// StartCodespace calls StartCodespaceFunc.
func (mock *apiClientMock) StartCodespace(ctx context.Context, name string) error {
	if mock.StartCodespaceFunc == nil {
		panic("apiClientMock.StartCodespaceFunc: method is nil but apiClient.StartCodespace was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockStartCodespace.Lock()
	mock.calls.StartCodespace = append(mock.calls.StartCodespace, callInfo)
	mock.lockStartCodespace.Unlock()
	return mock.StartCodespaceFunc(ctx, name)
}

// StartCodespaceCalls gets all the calls that were made to StartCodespace.
// Check the length with:
//
//	len(mockedapiClient.StartCodespaceCalls())
func (mock *apiClientMock) StartCodespaceCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockStartCodespace.RLock()
	calls = mock.calls.StartCodespace
	mock.lockStartCodespace.RUnlock()
	return calls
}

// StopCodespace calls StopCodespaceFunc.
func (mock *apiClientMock) StopCodespace(ctx context.Context, name string, orgName string, userName string) error {
	if mock.StopCodespaceFunc == nil {
		panic("apiClientMock.StopCodespaceFunc: method is nil but apiClient.StopCodespace was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Name     string
		OrgName  string
		UserName string
	}{
		Ctx:      ctx,
		Name:     name,
		OrgName:  orgName,
		UserName: userName,
	}
	mock.lockStopCodespace.Lock()
	mock.calls.StopCodespace = append(mock.calls.StopCodespace, callInfo)
	mock.lockStopCodespace.Unlock()
	return mock.StopCodespaceFunc(ctx, name, orgName, userName)
}

// StopCodespaceCalls gets all the calls that were made to StopCodespace.
// Check the length with:
//
//	len(mockedapiClient.StopCodespaceCalls())
func (mock *apiClientMock) StopCodespaceCalls() []struct {
	Ctx      context.Context
	Name     string
	OrgName  string
	UserName string
} {
	var calls []struct {
		Ctx      context.Context
		Name     string
		OrgName  string
		UserName string
	}
	mock.lockStopCodespace.RLock()
	calls = mock.calls.StopCodespace
	mock.lockStopCodespace.RUnlock()
	return calls
}
