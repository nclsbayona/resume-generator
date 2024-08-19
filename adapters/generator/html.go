package generator

// Options: template (Template directory to use)

import (
	"os"
	"text/template"

	"github.com/nclsbayona/resume-generator/domain"
	"github.com/nclsbayona/resume-generator/ports"
)

type AuxiliaryWriter struct {
	variable *string
}

func (a AuxiliaryWriter) Write(p []byte) (n int, err error) {
	tmp := *a.variable + string(p)
	*a.variable = tmp
	return len(p), nil
}

type AuxiliaryRender struct {
	Summary            *string
	Name               *string
	Experience         *string
	Education          *string
	ExtraInfo          *string
	ContinousEducation *string
}

func newAuxiliaryRender(name *string, summary *string) *AuxiliaryRender {
	return &AuxiliaryRender{Name: name, Summary: summary}
}

type WriteHTML struct {
	template *string
	resume   *domain.FullResume
}

func NewWriteHTML() *ports.Generator {
	var generator ports.Generator = &WriteHTML{}
	return &generator
}

func listTemplateDirectory(template_location *string) ([]*string, error) {
	list, err := os.ReadDir(*template_location)
	if err != nil {
		return nil, err
	}
	var files []*string
	for _, file := range list {
		fileName := file.Name()
		files = append(files, &fileName)
	}
	return files, nil
}

func renderExtraInfo(extra_info *domain.ExtraInformation, template_location *string, logger *domain.Logger) (rendered *string) {
	rendered = new(string)
	tmpl, err := template.ParseFiles(*template_location)
	if err != nil {
		logger.Warn("Error parsing template for experience " + err.Error())
	}
	var auxiliary_writer AuxiliaryWriter = AuxiliaryWriter{variable: rendered}
	logger.Info("Rendering " + *extra_info.Label + " with value " + *extra_info.Value)
	err = tmpl.Execute(auxiliary_writer, extra_info)
	if err != nil {
		logger.Warn("Error executing template for experience " + err.Error())
	}
	return
}

func renderExtraInfos(extra_infos []*domain.ExtraInformation, template *string, logger *domain.Logger) (rendered *string) {
	temporary_rendered := make([]*string, len(extra_infos))
	for i, info := range extra_infos {
		temporary_rendered[i] = renderExtraInfo(info, template, logger)
	}
	rendered = new(string)
	for _, rendered_educations := range temporary_rendered {
		tmp := *rendered + *rendered_educations + "\n"
		rendered = &tmp
	}
	return
}

func renderContinousEducation(continous_education *domain.ContinousEducation, template_location *string, logger *domain.Logger) (rendered *string) {
	rendered = new(string)
	tmpl, err := template.ParseFiles(*template_location)
	if err != nil {
		logger.Warn("Error parsing template for experience " + err.Error())
	}
	var auxiliary_writer AuxiliaryWriter = AuxiliaryWriter{variable: rendered}
	logger.Info("Rendering continous education " + *continous_education.Title + " at " + *continous_education.Institution + " gotten " + *continous_education.Date + " and valid to " + *continous_education.Expire + " with summary " + *continous_education.Summary)
	err = tmpl.Execute(auxiliary_writer, continous_education)
	if err != nil {
		logger.Warn("Error executing template for experience " + err.Error())
	}
	return
}

func renderContinousEducations(continous_educations []*domain.ContinousEducation, template *string, logger *domain.Logger) (rendered *string) {
	temporary_rendered := make([]*string, len(continous_educations))
	for i, info := range continous_educations {
		temporary_rendered[i] = renderContinousEducation(info, template, logger)
	}
	rendered = new(string)
	for _, rendered_educations := range temporary_rendered {
		tmp := *rendered + *rendered_educations + "\n"
		rendered = &tmp
	}
	return
}

func renderEducation(education *domain.Education, template_location *string, logger *domain.Logger) (rendered *string) {
	rendered = new(string)
	tmpl, err := template.ParseFiles(*template_location)
	if err != nil {
		logger.Warn("Error parsing template for experience " + err.Error())
	}
	var auxiliary_writer AuxiliaryWriter = AuxiliaryWriter{variable: rendered}
	logger.Info("Rendering experience " + *education.Title + " at " + *education.Institution + " from " + *education.StartDate + " to " + *education.EndDate + " with description " + *education.Description)
	err = tmpl.Execute(auxiliary_writer, education)
	if err != nil {
		logger.Warn("Error executing template for experience " + err.Error())
	}
	return
}

func renderEducations(educations []*domain.Education, template *string, logger *domain.Logger) (rendered *string) {
	temporary_rendered := make([]*string, len(educations))
	for i, education := range educations {
		temporary_rendered[i] = renderEducation(education, template, logger)
	}
	rendered = new(string)
	for _, rendered_educations := range temporary_rendered {
		tmp := *rendered + *rendered_educations + "\n"
		rendered = &tmp
	}
	return
}

func renderExperience(experience *domain.Experience, template_location *string, logger *domain.Logger) (rendered *string) {
	rendered = new(string)
	tmpl, err := template.ParseFiles(*template_location)
	if err != nil {
		logger.Warn("Error parsing template for experience " + err.Error())
	}
	var auxiliary_writer AuxiliaryWriter = AuxiliaryWriter{variable: rendered}
	logger.Info("Rendering experience " + *experience.Title + " at " + *experience.Company + " from " + *experience.StartDate + " to " + *experience.EndDate)
	err = tmpl.Execute(auxiliary_writer, experience)
	if err != nil {
		logger.Warn("Error executing template for experience " + err.Error())
	}
	return
}

func renderExperiences(experiences []*domain.Experience, template *string, logger *domain.Logger) (rendered *string) {
	temporary_rendered := make([]*string, len(experiences))
	for i, experience := range experiences {
		temporary_rendered[i] = renderExperience(experience, template, logger)
	}
	rendered = new(string)
	for _, rendered_experience := range temporary_rendered {
		tmp := *rendered + *rendered_experience + "\n"
		rendered = &tmp
	}
	return
}

func sliceContains(slice []*string, search *string) (i int, ok bool) {
	i = -1
	for j, s := range slice {
		if *s == *search {
			ok = true
			i = j
		}
	}
	return
}

func (w *WriteHTML) GetFull(logger *domain.Logger) (rendered *string) {
	logger.Info("Using template: " + *w.template + " for resume of: " + *w.resume.Name)
	template_contents, err := listTemplateDirectory(w.template)
	if err != nil {
		logger.Warn("Error listing template directory: " + *w.template)
		logger.Error(err.Error())
	}
	if len(template_contents) == 0 {
		logger.Error("No templates found for template: " + *w.template)
	}

	var auxiliary_renderable *AuxiliaryRender = newAuxiliaryRender(w.resume.Name, w.resume.Summary)
	// Render experiences
	string_to_search := "experience.html.tpl"
	if index, ok := sliceContains(template_contents, &string_to_search); ok {
		logger.Info("Rendering experience with template " + *template_contents[index])
		tmp_experiences := *w.template + "/" + *template_contents[index]
		auxiliary_renderable.Experience = renderExperiences(w.resume.Experiences, &tmp_experiences, logger)
	} else {
		var tmp_experiences string = ""
		auxiliary_renderable.Experience = &tmp_experiences
		logger.Warn("No template for experience found")
	}

	// Render education
	string_to_search = "education.html.tpl"
	if index, ok := sliceContains(template_contents, &string_to_search); ok {
		logger.Info("Rendering education with template " + *template_contents[index])
		tmp_education := *w.template + "/" + *template_contents[index]
		auxiliary_renderable.Education = renderEducations(w.resume.Education, &tmp_education, logger)
	} else {
		var tmp_education string = ""
		auxiliary_renderable.Education = &tmp_education
		logger.Warn("No template for education found")
	}

	// Render extra info
	string_to_search = "extra.html.tpl"
	if index, ok := sliceContains(template_contents, &string_to_search); ok {
		logger.Info("Rendering extra info with template " + *template_contents[index])
		tmp_extra := *w.template + "/" + *template_contents[index]
		auxiliary_renderable.ExtraInfo = renderExtraInfos(w.resume.ExtraInfo, &tmp_extra, logger)
	} else {
		var tmp_extra string = ""
		auxiliary_renderable.ExtraInfo = &tmp_extra
		logger.Warn("No template for extra info found")
	}

	// Render continous education
	string_to_search = "continous_education.html.tpl"
	if index, ok := sliceContains(template_contents, &string_to_search); ok {
		logger.Info("Rendering continous education with template " + *template_contents[index])
		tmp_continous_education := *w.template + "/" + *template_contents[index]
		auxiliary_renderable.ContinousEducation = renderContinousEducations(w.resume.ContinousEducation, &tmp_continous_education, logger)
	} else {
		var tmp_continous_education string = ""
		auxiliary_renderable.ContinousEducation = &tmp_continous_education
		logger.Warn("No template for continous education found")
	}

	// Render full resume
	string_to_search = "full.html.tpl"
	index, ok := sliceContains(template_contents, &string_to_search)
	var full_resume_template_path *string
	if !ok {
		logger.Warn("No template for full resume found")
		tmp := "templates/basic/full.html.tpl"
		full_resume_template_path = &tmp
	} else {
		full_resume_template_path = template_contents[index]
	}
	tmp := *w.template + "/" + *full_resume_template_path
	tmpl, err := template.ParseFiles(tmp)
	if err != nil {
		logger.Error(err.Error())
	}
	aux := ""
	auxiliary_writer := AuxiliaryWriter{variable: &aux}
	err = tmpl.Execute(auxiliary_writer, auxiliary_renderable)
	rendered = &aux
	if err != nil {
		logger.Error(err.Error())
	}
	return
}

func (w *WriteHTML) SetOptions(options ...*string) {
	w.template = options[0]
}

func (w *WriteHTML) SetResume(resume *domain.FullResume) {
	w.resume = resume
}
