#!/bin/bash
npm install
if [ ! -d $(pwd)/public/js/angular/lib/ ]; then
  mkdir $(pwd)/public/js/angular/lib/
fi
cp -r $(pwd)/node_modules/angular2 $(pwd)/public/js/angular/lib/
cp -r $(pwd)/node_modules/rxjs $(pwd)/public/js/angular/lib/
