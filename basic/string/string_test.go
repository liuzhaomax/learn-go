package main

import (
	"fmt"
	"testing"
)

const InitUserInfo = `
{
  "basicInfo": {
    "avatar": "",
    "nickname": "",
    "name": {
      "firstName": "",
      "lastName": ""
    },
    "occupation": "",
    "company": "",
    "degree": "",
    "birthday": "",
    "location": "",
    "contactInfo": {
      "email": "",
      "phoneNumber": "",
      "officeAddress": ""
    },
    "personalIntroduction": ""
  },
  "careerHistory": {
    "academicActivities": [],
    "educationalExperience": [],
    "awards": [],
    "teachingExperience": [],
    "workExperience": [],
    "relatedPersons": []
  },
  "publications": {
    "books": [],
    "articles": [],
    "papers": [],
    "codeProjects": []
  },
  "relatedLinks": []
}
`

func TestStringFanQuote(t *testing.T) {
	fmt.Println(InitUserInfo)
}
