<li>
<div>
<h3>{{.Title}}</h3>        
<div>
<p>{{.Company}} at {{.Location}}</p>     
<p>{{.StartDate}} to {{.EndDate}}</p>
</div>
<span>{{.Description}}</span>
<b>Achievements</b>
<ul>
{{range .Achievements}}<li>{{.}}</li>
{{end}}
</ul>
</div>
</li>