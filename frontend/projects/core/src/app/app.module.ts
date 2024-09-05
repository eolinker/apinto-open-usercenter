import {  NgModule } from '@angular/core';
import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { LayoutModule } from './layout/layout.module';
import { EoNgTableModule } from 'eo-ng-table';
import { EoNgFeedbackMessageService } from 'eo-ng-feedback';
import { HttpClientModule } from '@angular/common/http';
import { NZ_CONFIG, NzConfig } from 'ng-zorro-antd/core/config';
import { ApiService, } from './service/api.service';
import { ModuleFederationService } from './service/module-federation.service';
import { CommonModule } from '@angular/common';

const ngZorroConfig: NzConfig = {
  // 注意组件名称没有 nz 前缀
  message: { nzMaxStack: 1, nzDuration: 2000 },
  notification: { nzMaxStack: 1, nzDuration: 2000 }
}

declare global{
  interface Window{
    apinto?:any
    apintoDebug?:any
    _apinto_mf?:boolean
  }
}

let coreUserService:any

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    CommonModule,
    EoNgTableModule,
    LayoutModule,
    HttpClientModule,
    AppRoutingModule,
  ],
  providers: [
    ApiService,
    EoNgFeedbackMessageService,
    { provide: NZ_CONFIG, useValue: ngZorroConfig },
    ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(private moduleFederationService:ModuleFederationService){
    if(!this.moduleFederationService.initialized){
      this.moduleFederationService.providerFromCore = coreUserService.pluginProvider
      this.moduleFederationService.pluginSlotHub = coreUserService.pluginSlotHub
      this.moduleFederationService.initialized = true
    } 
  }

 }

export function bootstrap (props:any){ 
  coreUserService  = props
}
export async function beforeMount  (props:any){
  coreUserService.pluginProvider.redirectUrl  = props.redirectUrl
};
export async function mount  (props:any){
};
export function beforeUnmount  (props:any){ }
export function unmount  (props:any){ }
