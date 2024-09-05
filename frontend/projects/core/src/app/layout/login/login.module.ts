import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login.component'
import { LoginRoutingModule } from './login-routing.module';
import { LogoComponent } from './logo/logo.component'
import { PasswordComponent } from './password/password.component'
import { EoNgButtonModule } from 'eo-ng-button';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { NzFormModule } from 'ng-zorro-antd/form';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { EoNgFeedbackMessageModule } from 'eo-ng-feedback';
import { ApiService } from '../../service/api.service';
import { EoNgInputModule } from 'eo-ng-input';
import { ModuleFederationService } from '../../service/module-federation.service';
let coreUserService:any

@NgModule({
  declarations: [
    LoginComponent,
    LogoComponent,
    PasswordComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    LoginRoutingModule,
    EoNgButtonModule,
    EoNgInputModule,
    EoNgFeedbackMessageModule,
    NzGridModule,
    NzFormModule,
    
  ],
  providers: [
    ApiService,
    // { provide: API_URL, useValue: environment.urlPrefix },
    ],
  })
export class LoginModule {
  constructor(private moduleFederationService:ModuleFederationService){
    if(!this.moduleFederationService.initialized){
      this.moduleFederationService.providerFromCore = coreUserService.pluginProvider
      this.moduleFederationService.pluginSlotHub = coreUserService.pluginSlotHub
      this.moduleFederationService.initialized = true
    }
  } }


export function bootstrap (props:any){ 
  coreUserService  = props}
export async function beforeMount  (props:any){
};
export async function mount  (props:any){
};
export function beforeUnmount  (props:any){}
export function unmount  (props:any){}