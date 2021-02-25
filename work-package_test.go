package openproject

import (
	"fmt"
	"net/http"
	"testing"
)

const testWorkPackageGETResponse = `{"derivedStartDate":null,"derivedDueDate":null,"_embedded":{"attachments":{"_type":"Collection","total":0,"count":0,"_embedded":{"elements":[]},"_links":{"self":{"href":"/api/v3/work_packages/36350/attachments"}}},"relations":{"_type":"Collection","total":0,"count":0,"_embedded":{"elements":[]},"_links":{"self":{"href":"/api/v3/work_packages/36350/relations"}}},"type":{"_type":"Type","id":1,"name":"Bug","color":"#FF0013","position":4,"isDefault":false,"isMilestone":false,"createdAt":"2013-11-13T13:11:21Z","updatedAt":"2020-11-25T13:57:59Z","_links":{"self":{"href":"/api/v3/types/1","title":"Bug"}}},"priority":{"_type":"Priority","id":4,"name":"Medium","position":2,"color":"#51CF66","isDefault":true,"isActive":true,"_links":{"self":{"href":"/api/v3/priorities/4","title":"Medium"}}},"project":{"_type":"Project","id":14,"identifier":"openproject","name":"OpenProject","active":true,"public":true,"description":{"format":"markdown","raw":"Building the best open source project collaboration software.","html":"<p class=\"op-uc-p\">Building the best open source project collaboration software.</p>"},"createdAt":"2012-09-21T13:37:39Z","updatedAt":"2021-02-16T12:50:23Z","status":null,"statusExplanation":{"format":"markdown","raw":null,"html":""},"_links":{"self":{"href":"/api/v3/projects/14","title":"OpenProject"},"workPackages":{"href":"/api/v3/projects/14/work_packages"},"categories":{"href":"/api/v3/projects/14/categories"},"versions":{"href":"/api/v3/projects/14/versions"},"types":{"href":"/api/v3/projects/14/types"},"schema":{"href":"/api/v3/projects/schema"},"customField105":{"href":null,"title":null},"parent":{"href":null,"title":null}}},"status":{"_type":"Status","id":63,"name":"prioritized","isClosed":false,"color":"#D90000","isDefault":false,"isReadonly":false,"defaultDoneRatio":null,"position":24,"_links":{"self":{"href":"/api/v3/statuses/63","title":"prioritized"}}},"author":{"_type":"User","id":67841,"name":"Matthias Laux","avatar":"https://secure.gravatar.com/avatar/c2b652a24efa0e54eff27b1b70e58518?default=404&secure=true","_links":{"self":{"href":"/api/v3/users/67841","title":"Matthias Laux"},"showUser":{"href":"/users/67841","type":"text/html"}}},"customActions":[]},"_type":"WorkPackage","id":36350,"lockVersion":3,"subject":"Default work package type can't be changed","description":{"format":"markdown","raw":"### **Steps to reproduce:**\n\n1.  Go to /types\n2.  Move a random type to the top of the list, using the arrows\n3.  Go to a project&#39;s work package list and create a new work package, using the shortcut (type &quot;nwp&quot;) (or create a new work package using the &quot;+ Create new work package link&quot;)\n\n### **Actual Behavior**\n\nThe new work package is still of the type task\n\n### **Expected Behavior**\n\nThe work package type I moved to the top of the list is the current default type and will be selected when I create a new work package.","html":"<h3 id=\"steps-to-reproduce\" class=\"op-uc-h3\">\n<strong>Steps to reproduce:</strong><a class=\"op-uc-link_permalink icon-link op-uc-link\" aria-hidden=\"true\" href=\"#steps-to-reproduce\"></a>\n</h3>\n<ol class=\"op-uc-list\">\n<li class=\"op-uc-list--item\">Go to /types</li>\n<li class=\"op-uc-list--item\">Move a random type to the top of the list, using the arrows</li>\n<li class=\"op-uc-list--item\">Go to a project's work package list and create a new work package, using the shortcut (type \"nwp\") (or create a new work package using the \"+ Create new work package link\")</li>\n</ol>\n<h3 id=\"actual-behavior\" class=\"op-uc-h3\">\n<strong>Actual Behavior</strong><a class=\"op-uc-link_permalink icon-link op-uc-link\" aria-hidden=\"true\" href=\"#actual-behavior\"></a>\n</h3>\n<p class=\"op-uc-p\">The new work package is still of the type task</p>\n<h3 id=\"expected-behavior\" class=\"op-uc-h3\">\n<strong>Expected Behavior</strong><a class=\"op-uc-link_permalink icon-link op-uc-link\" aria-hidden=\"true\" href=\"#expected-behavior\"></a>\n</h3>\n<p class=\"op-uc-p\">The work package type I moved to the top of the list is the current default type and will be selected when I create a new work package.</p>"},"scheduleManually":false,"startDate":null,"dueDate":null,"estimatedTime":null,"derivedEstimatedTime":null,"createdAt":"2021-02-23T11:09:23Z","updatedAt":"2021-02-23T11:21:27Z","position":1,"storyPoints":null,"remainingTime":null,"customField108":true,"customField130":null,"customField114":null,"_links":{"attachments":{"href":"/api/v3/work_packages/36350/attachments"},"prepareAttachment":{"href":"/api/v3/work_packages/36350/attachments/prepare","method":"post"},"self":{"href":"/api/v3/work_packages/36350","title":"Default work package type can't be changed"},"schema":{"href":"/api/v3/work_packages/schemas/14-1"},"availableRelationCandidates":{"href":"/api/v3/work_packages/36350/available_relation_candidates","title":"Potential work packages to relate to"},"activities":{"href":"/api/v3/work_packages/36350/activities"},"relations":{"href":"/api/v3/work_packages/36350/relations"},"revisions":{"href":"/api/v3/work_packages/36350/revisions"},"previewMarkup":{"href":"/api/v3/render/markdown?context=/api/v3/work_packages/36350","method":"post"},"category":{"href":null},"type":{"href":"/api/v3/types/1","title":"Bug"},"priority":{"href":"/api/v3/priorities/4","title":"Medium"},"project":{"href":"/api/v3/projects/14","title":"OpenProject"},"status":{"href":"/api/v3/statuses/63","title":"prioritized"},"author":{"href":"/api/v3/users/67841","title":"Matthias Laux"},"responsible":{"href":null},"assignee":{"href":null},"version":{"href":null},"customField1":[{"title":"Work packages","href":"/api/v3/custom_options/32"},{"title":"Administration","href":"/api/v3/custom_options/396"}],"customField5":{"href":null,"title":null},"customField107":[],"customField102":{"href":null,"title":null},"customField127":[],"customField129":{"href":null,"title":null},"ancestors":[],"parent":{"href":null,"title":null},"customActions":[]}}`
const testWorkPackagePOSTResponse = `{"derivedStartDate":null,"derivedDueDate":null,"_embedded":{"attachments":{"_type":"Collection","total":0,"count":0,"_embedded":{"elements":[]},"_links":{"self":{"href":"/api/v3/work_packages/38/attachments"}}},"relations":{"_type":"Collection","total":0,"count":0,"_embedded":{"elements":[]},"_links":{"self":{"href":"/api/v3/work_packages/38/relations"}}},"type":{"_type":"Type","id":1,"name":"Task","color":"#1A67A3","position":1,"isDefault":true,"isMilestone":false,"createdAt":"2021-02-18T10:10:30Z","updatedAt":"2021-02-18T10:10:30Z","_links":{"self":{"href":"/api/v3/types/1","title":"Task"}}},"priority":{"_type":"Priority","id":8,"name":"Normal","position":2,"color":"#74C0FC","isDefault":true,"isActive":true,"_links":{"self":{"href":"/api/v3/priorities/8","title":"Normal"}}},"project":{"_type":"Project","id":1,"identifier":"demo-project","name":"Demo project","active":true,"public":true,"description":{"format":"markdown","raw":"This is a short summary of the goals of this demo project.","html":"<p class=\"op-uc-p\">This is a short summary of the goals of this demo project.</p>"},"createdAt":"2021-02-18T10:10:47Z","updatedAt":"2021-02-18T10:10:47Z","status":"on track","statusExplanation":{"format":"markdown","raw":"All tasks are on schedule. The people involved know their tasks. The system is completely set up.","html":"<p class=\"op-uc-p\">All tasks are on schedule. The people involved know their tasks. The system is completely set up.</p>"},"_links":{"self":{"href":"/api/v3/projects/1","title":"Demo project"},"createWorkPackage":{"href":"/api/v3/projects/1/work_packages/form","method":"post"},"createWorkPackageImmediately":{"href":"/api/v3/projects/1/work_packages","method":"post"},"workPackages":{"href":"/api/v3/projects/1/work_packages"},"categories":{"href":"/api/v3/projects/1/categories"},"versions":{"href":"/api/v3/projects/1/versions"},"memberships":{"href":"/api/v3/memberships?filters=%5B%7B%22project%22%3A%7B%22operator%22%3A%22%3D%22%2C%22values%22%3A%5B%221%22%5D%7D%7D%5D"},"types":{"href":"/api/v3/projects/1/types"},"update":{"href":"/api/v3/projects/1/form","method":"post"},"updateImmediately":{"href":"/api/v3/projects/1","method":"patch"},"delete":{"href":"/api/v3/projects/1","method":"delete"},"schema":{"href":"/api/v3/projects/schema"},"parent":{"href":null,"title":null}}},"status":{"_type":"Status","id":1,"name":"New","isClosed":false,"color":"#1098AD","isDefault":true,"isReadonly":false,"defaultDoneRatio":null,"position":1,"_links":{"self":{"href":"/api/v3/statuses/1","title":"New"}}},"author":{"_type":"User","id":1,"name":"Manuelbcd","createdAt":"2021-02-18T10:10:46Z","updatedAt":"2021-02-23T09:17:44Z","login":"user@acme.com","admin":true,"firstName":"Manuel","lastName":"Bcd","email":"user@acme.com","avatar":"https://secure.gravatar.com/avatar/af5c1e5d637c5bdf7be72b0904272a89?default=404&secure=true","status":"active","identityUrl":null,"language":"en","_links":{"self":{"href":"/api/v3/users/1","title":"Manuelbcd"},"memberships":{"href":"/api/v3/memberships?filters=%5B%7B%22principal%22%3A%7B%22operator%22%3A%22%3D%22%2C%22values%22%3A%5B%221%22%5D%7D%7D%5D","title":"Members"},"showUser":{"href":"/users/1","type":"text/html"},"updateImmediately":{"href":"/api/v3/users/1","title":"Update user@acme.com","method":"patch"},"lock":{"href":"/api/v3/users/1/lock","title":"Set lock on user@acme.com","method":"post"}}},"customActions":[]},"_type":"WorkPackage","id":38,"lockVersion":0,"subject":"subject34","description":{"format":"markdown","raw":"esto es una prueba","html":"<p class=\"op-uc-p\">esto es una prueba</p>"},"scheduleManually":false,"startDate":null,"dueDate":null,"estimatedTime":null,"derivedEstimatedTime":null,"percentageDone":0,"createdAt":"2021-02-23T15:11:35Z","updatedAt":"2021-02-23T15:11:35Z","_links":{"attachments":{"href":"/api/v3/work_packages/38/attachments"},"prepareAttachment":{"href":"/api/v3/work_packages/38/attachments/prepare","method":"post"},"addAttachment":{"href":"/api/v3/work_packages/38/attachments","method":"post"},"self":{"href":"/api/v3/work_packages/38","title":"subject34"},"update":{"href":"/api/v3/work_packages/38/form","method":"post"},"schema":{"href":"/api/v3/work_packages/schemas/1-1"},"updateImmediately":{"href":"/api/v3/work_packages/38","method":"patch"},"delete":{"href":"/api/v3/work_packages/38","method":"delete"},"move":{"href":"/work_packages/38/move/new","type":"text/html","title":"Move subject34"},"copy":{"href":"/work_packages/38/copy","title":"Copy subject34"},"pdf":{"href":"/work_packages/38.pdf","type":"application/pdf","title":"Export as PDF"},"atom":{"href":"/work_packages/38.atom","type":"application/rss+xml","title":"Atom feed"},"availableRelationCandidates":{"href":"/api/v3/work_packages/38/available_relation_candidates","title":"Potential work packages to relate to"},"customFields":{"href":"/projects/demo-project/settings/custom_fields","type":"text/html","title":"Custom fields"},"configureForm":{"href":"/types/1/edit?tab=form_configuration","type":"text/html","title":"Configure form"},"activities":{"href":"/api/v3/work_packages/38/activities"},"availableWatchers":{"href":"/api/v3/work_packages/38/available_watchers"},"relations":{"href":"/api/v3/work_packages/38/relations"},"revisions":{"href":"/api/v3/work_packages/38/revisions"},"watchers":{"href":"/api/v3/work_packages/38/watchers"},"addWatcher":{"href":"/api/v3/work_packages/38/watchers","method":"post","payload":{"user":{"href":"/api/v3/users/{user_id}"}},"templated":true},"removeWatcher":{"href":"/api/v3/work_packages/38/watchers/{user_id}","method":"delete","templated":true},"addRelation":{"href":"/api/v3/work_packages/38/relations","method":"post","title":"Add relation"},"addChild":{"href":"/api/v3/projects/demo-project/work_packages","method":"post","title":"Add child of subject34"},"changeParent":{"href":"/api/v3/work_packages/38","method":"patch","title":"Change parent of subject34"},"addComment":{"href":"/api/v3/work_packages/38/activities","method":"post","title":"Add comment"},"previewMarkup":{"href":"/api/v3/render/markdown?context=/api/v3/work_packages/38","method":"post"},"category":{"href":null},"type":{"href":"/api/v3/types/1","title":"Task"},"priority":{"href":"/api/v3/priorities/8","title":"Normal"},"project":{"href":"/api/v3/projects/1","title":"Demo project"},"status":{"href":"/api/v3/statuses/1","title":"New"},"author":{"href":"/api/v3/users/1","title":"Manuelbcd"},"responsible":{"href":null},"assignee":{"href":null},"version":{"href":null},"watch":{"href":"/api/v3/work_packages/38/watchers","method":"post","payload":{"user":{"href":"/api/v3/users/1"}}},"ancestors":[],"parent":{"href":null,"title":null},"customActions":[]}}`

func TestWorkPackageService_Get_Success(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/api/v3/work_packages/36350", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/api/v3/work_packages/36350")

		fmt.Fprint(w, testWorkPackageGETResponse)
	})

	workpkg, _, err := testClient.WorkPackage.Get("36350", nil)
	if workpkg == nil {
		t.Error("Expected work-package. Work-package is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestWorkPackageService_Get_SearchListSuccess(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/api/v3/work_packages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/api/v3/work_packages")

		fmt.Fprint(w, testWorkPackageGETResponse)
	})

	workpkg, _, err := testClient.WorkPackage.Get("36350", nil)
	if workpkg == nil {
		t.Error("Expected work-package. Work-package is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestWorkPackageService_Create(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/api/v3/projects/demo-project/work_packages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/api/v3/projects/demo-project/work_packages")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, testWorkPackagePOSTResponse)
	})

	i := &WorkPackage{
		Subject: "Just another test work-package",
		Description: &WPDescription{
			Format: "textile",
			Raw:    "This is just a demo work-package description",
		},
	}
	issue, _, err := testClient.WorkPackage.Create(i, "demo-project")
	if issue == nil {
		t.Error("Expected work-package. Work-package is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
