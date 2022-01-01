//go:generate esbuild app.jsx --bundle --minify --sourcemap --outfile=build/app.js
//go:generate cp -r public/ build/

package apidocs
