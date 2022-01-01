//go:generate yarn
//go:generate cp -r public/ build/
//go:generate esbuild app.jsx --bundle --minify --sourcemap --outfile=build/app.js --define:global=window --inject:esbuild.inject.js

package apidocs
