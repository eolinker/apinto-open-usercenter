import { environment } from "projects/core/src/environments/environment";

const ModuleFederationPlugin = require("webpack/lib/container/ModuleFederationPlugin");
const mf = require("@angular-architects/module-federation/webpack");
const path = require("path");
const share = mf.share;
const { dependencies, devDependencies } = require('./package.json')

const sharedMappings = new mf.SharedMappings();
sharedMappings.register(
  path.join(__dirname, './tsconfig.json'),
  [/* mapped paths to share */]);

module.exports = {
  output: {
    uniqueName: environment.pluginName,
    publicPath: `/plugin-frontend/${environment.pluginName}/`,
  },
  optimization: {
    runtimeChunk: false
  },   
  resolve: {
    alias: {
      ...sharedMappings.getAliases(),
    }
  },
  experiments: {
    outputModule: true
  },
  plugins: [
    new ModuleFederationPlugin({
        library: { type: "module" },
        name: environment.pluginName,
        filename: "apinto.js",
        exposes: {
            './AppModule': './projects/core/src/app/app.module.ts',
            './LoginModule': './projects/core/src/app/layout/login/login.module.ts',
            './Bootstrap':'./projects/core/src/app/layout/bootstrap/bootstrap.ts'
          },        
        // For hosts (please adjust)
        // remotes: {
        //     "main": "http://localhost:9000/remoteEntry.js",
        // },
        shared: share({
          "@angular/core": { singleton: true, strictVersion: true, requiredVersion: 'auto' }, 
          "@angular/common": { singleton: true, strictVersion: true, requiredVersion: 'auto' }, 
          "@angular/common/http": { singleton: true, strictVersion: true, requiredVersion: 'auto' }, 
          "@angular/router": { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          "@angular/animations": { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          "@angular/platform-browser/animations": { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          '@angular/forms': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          "eo-ng-feedback":{ singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'crypto-js': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'lodash-es': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'ng-zorro-antd': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-select': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-auto-complete': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-breadcrumb': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-button': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-cascader': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-checkbox': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-copy': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-dropdown': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-empty': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-input': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-layout': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-menu': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-radio': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-switch': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-table': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          'eo-ng-tree': { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          rxjs: { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          tailwindcss: { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          uuid: { singleton: true, strictVersion: true, requiredVersion: 'auto' },
          ...sharedMappings.getDescriptors()
        })
    }),
    sharedMappings.getPlugin()
  ],
};
