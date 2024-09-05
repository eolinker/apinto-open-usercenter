import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { EoNgDropdownModule } from 'eo-ng-dropdown';
import { UserAvatarComponent } from './user-avatar.component';
import { BrowserModule } from '@angular/platform-browser';
import { UserProfileModule } from '../user-profile/user-profile.module';
import { EoNgFeedbackMessageService } from 'eo-ng-feedback';
import { ModuleFederationService } from '../../service/module-federation.service';

let coreUserService:any

@NgModule({
  declarations: [UserAvatarComponent
  ],
  imports: [
    CommonModule,
    BrowserModule,
    EoNgDropdownModule,
    UserProfileModule
  ],
  exports:[UserAvatarComponent
  ],
  providers:[ EoNgFeedbackMessageService 
    ]
})
export class UserAvatarModule { 
  constructor(private moduleFederationService:ModuleFederationService){
    if(!this.moduleFederationService.providerFromCore){
      this.moduleFederationService.providerFromCore = coreUserService
    } 
  }
  ngDoBootstrap() {}
}

export function bootstrap( props: any): void {
  const { pluginEventHub, pluginSlotHub, pluginProvider,injector,closeModal,router,messageService,modalService} = props
  coreUserService  = pluginProvider
}
