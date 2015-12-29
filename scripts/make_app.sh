#!/bin/bash
npm install
cp -r node_modules/angular2 public/js/angular/lib/
cp -r node_modules/systemjs public/js/angular/lib/
cp -r node_modules/rxjs public/js/angular/lib/
grunt browserify
