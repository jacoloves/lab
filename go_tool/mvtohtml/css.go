package main

func css() string {
	return `
	<style>
	body {
		font-family: -apple-system, BlinkMacSystemFont, "Segoe WPC", "Segoe UI", "Ubuntu", "Droid Sans", sans-serif, "Meiryo";
		padding: 0 12px;
	}

	img {
		max-width: 100%;
		max-height: 100%;
	}
	
	a {
		text-decoration: none;
	}
	
	a:hover {
		text-decoration: underline;
	}
	
	a:focus,
	input:focus,
	select:focus,
	textarea:focus {
		outline: 1px solid -webkit-focus-ring-color;
		outline-offset: -1px;
	}
	
	hr {
		border: 0;
		height: 2px;
		border-bottom: 2px solid;
	}
	
	h1 {
		padding-bottom: 0.3em;
		line-height: 1.2;
		border-bottom-width: 1px;
		border-bottom-style: solid;
	}
	
	h1, h2, h3 {
		font-weight: normal;
	}
	
	table {
		border-collapse: collapse;
	}
	
	table > thead > tr > th {
		text-align: left;
		border-bottom: 1px solid;
	}
	
	table > thead > tr > th,
	table > thead > tr > td,
	table > tbody > tr > th,
	table > tbody > tr > td {
		padding: 5px 10px;
	}
	
	table > tbody > tr + tr > td {
		border-top: 1px solid;
	}
	
	blockquote {
		margin: 0 7px 0 5px;
		padding: 0 16px 0 10px;
		border-left-width: 5px;
		border-left-style: solid;
	}
	
	code {
		font-family: Menlo, Monaco, Consolas, "Droid Sans Mono", "Courier New", monospace, "Droid Sans Fallback";
		font-size: 1em;
		line-height: 1.357em;
	}
	
	body.wordWrap pre {
		white-space: pre-wrap;
	}
	
	pre:not(.hljs),
	pre.hljs code > div {
		padding: 16px;
		border-radius: 3px;
		overflow: auto;
	}
	
	pre code {
		color: var(--vscode-editor-foreground);
		tab-size: 4;
	}
</style>
`
}
