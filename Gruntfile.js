module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    browserify: {
      main: {
        src: 'entry.js',
        dest: 'public/js/angular/bundle.js'
      }
    }
  });

  grunt.loadNpmTasks('grunt-browserify');
};
