{{ define "index" }}

{{ template "header" }}




<div class="menu">
	<div class="container-fluid h1">
		Приветствуем Вас на сайте нашей строительной компании!
	</div>
	<div class="container-fluid h2">
		Мы поможем обустроить вашу квартиру, комнату, частный дом или любое другое помещение(нужное подчеркнуть) по вашему дизайну.
	</div>
	<div class="container-fluid h2">
		Или же мы предложим дизайнера, который учтет все ваши предпочтения, напиши клиентами были такие люди как Иван Иванов и Петр Петров. Поверьте, они остались довольны полученным результатом!
	</div>
</div>



{{ range $key, $value := . }}
<div class="row">
	<div class="col-xs-2">
	</div>
	<div class="col-xs-8">
		<h1><a href="/edit?id={{$key}}">{{ $value.Title }}</a></h1>
		
	</div>
	<div class="col-xs-2">
	</div>
</div>
<div class="row">
	<div class="col-xs-2">
	</div>
	<div class="col-xs-8">
		{{ $value.Content }}
	</div>
	<div class="col-xs-2">
	</div>
</div>
{{ end }}

{{ template "footer" }}

{{ end }}