mkdir -p ./build
sass ./css/misc.scss ./css/styles.css
cp -r -v ./css/*.css ./build
go run .
