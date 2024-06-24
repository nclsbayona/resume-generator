<html>
<head>
</head>
<body>
<h1>{{.Name}}</h1>
<h2?{{.Summary}}</h2>
<section id="extra-section">
<h2>Extra Information</h2>
<ul>
{{.ExtraInfo}}
</ul>
</section>
<section id="experiences-section">
<h3>Experience</h3>
<ul>
{{.Experience}}
</ul>
</section>
<section id="education-section">
<h3>Education</h3>
<ul>
{{.Education}}
</ul>
</section>
<section id="continous-section">
<h3>Continous Education</h3>
<ul>
{{.ContinousEducation}}
</ul>
</section>
</body>
</html>