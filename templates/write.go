{{ define "write" }}

{{ template "header" }}

<div class="py-5">
	<div class="container">
	<p></p>
	</div>
	<div class="container">
		<form role="form" method="POST" action="/SavePost">
			<input type="hidden" name="id" value="{{.Id}}" />
			<div class="container">
				<label><h3>Наше ФИО</h3></label>
				<input type="text" class="form-control" id="title" name="title" value="{{.Title}}"  />
			</div>
			<div class="container">
				<label><h3>Краткое описание работы, которую нужно будет проделать и контакты</h3></label>
				<input type="text" class="form-control" id="content" name="content" value="{{.Content}}"  />

			</div>
			<p></p>
			<button type="submit" style="margin-left: 45%;" class="btn btn-default">Отправить</button>
		</form>

	</div>
	<div class="col-xs-4">
		{{ if .Id }}
			<a href="/delete?id={{.Id}}">Delete</a>
		{{ end }}
	</div>
</div>

{{ template "footer" }}

{{ end }}