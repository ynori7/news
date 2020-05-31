package templates

const CoronaNewsTemplate = `<html>
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<style>
		h1 {
			text-align: center;
			margin-bottom: 5rem;
		}
		.info {
			font-size: .875rem;
    		line-height: 1.4;
    		color: #787878;
    		text-transform: uppercase;
		}
	</style>
</head>
<body>
<section class="news">
	<h1>Corona Nachrichten</h1>

	{{ range $i, $val := .News }}
	<div class="news-item">
		<h3>{{ $val.Title }}</h3>
		<div class="info">
			<span class="date">{{ $val.Date }}</span>
		</div>
		<div class="description"><p>{{ $val.Description }}</p></div>
	</div>
	<hr />
	{{ end }}
</section>
</body>
</html>
`

