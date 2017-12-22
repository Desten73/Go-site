{{ define "write" }}

{{ template "header" }}

<div class="row">
	<div class="col-xs-4">
	</div>
	<div class="col-xs-4">
		<form role="form" method="POST" action="/SavePost">
			<input type="hidden" name="id" value="{{.Id}}" />
			<div class="form-group">
				<label>Наше ФИО</label>
				<input type="text" class="form-control" id="title" name="title" value="{{.Title}}"  />
			</div>
			<div class="form-group">
				<label>Краткое описание работы, которую нужно будет проделать и контакты</label>
				<textarea id="content" name="content">{{.Content}}</textarea>
			</div>
			<button type="submit" class="btn btn-default">Отправить</button>
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