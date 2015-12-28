module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    wiredep: {
    task: {
      directory: 'public/bower_components',
      src: [
        'home.html'
      ]
    }
  }
  });

  grunt.loadNpmTasks('grunt-wiredep');
};
