package godebug

var debugger_css = `

	html {
	  position: relative;
	  min-height: 100%;
	}
	body {
	  margin-bottom: 300px;
	}
	.footer {
	  position: absolute;
	  bottom: 0;
	  width: 100%;
	  height: 300px;
	  background-color: #f5f5f5;
	}

	body > .container {
	  padding: 60px 15px 0;
	  overflow-y: auto;
	}
	.container .text-muted {
	  margin: 20px 0;
	}

	.nav li > a:hover {
		background-color: #e7e7e7;
	}

	.footer > .container {
	  padding-right: 15px;
	  padding-left: 15px;
	  overflow-y: auto;
	}

	code {
	  font-size: 80%;
	}

	.highlight {
		background-color: #ff2 !important;
	}

	.highlight.disabled {
	  text-decoration: line-through;
	}

	.navbar-brand {
		height: 60px;
		margin-left: 0 !important;
		line-height: 30px !important;
	}
`
