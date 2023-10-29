package contact


// Field names should start with an uppercase letter
type ContactFormDetails struct {
    Name string `json:"name" xml:"name" form:"name"`
    Email string `json:"email" xml:"email" form:"email"`
    Message string `json:"message" xml:"message" form:"message"`
}


// please see https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34