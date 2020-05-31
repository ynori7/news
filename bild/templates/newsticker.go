package templates

const NewsTickerTemplate = `<html>
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<style>
		h1 {
			text-align: center;
			margin-bottom: 3rem;
		}
		.info {
			font-size: .875rem;
    		line-height: 1.4;
    		color: #787878;
    		text-transform: uppercase;
		}
		.back {
			margin-bottom: 2rem;
		}
	</style>
</head>
<body>
<section class="news">
	<h1>Nachrichten</h1>

	<div class="back"><a href="/bild">Züruck</a></div>

	{{ range $i, $val := .News }}
	<div class="news-item">
		<h3><a href="{{ $val.Link }}" target="_blank">{{ $val.Title }}</a></h3>
		<div class="info">
			<span class="date">{{ $val.Date }}</span>
			<span class="category">{{ $val.Category }}</span>
		</div>
		<div class="description"><p>{{ $val.Description }}</p></div>
	</div>
	<hr />
	{{ end }}
</section>
</body>
</html>
`

