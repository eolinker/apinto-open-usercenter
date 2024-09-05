import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ResetPswComponent } from '../reset-psw/reset-psw.component';
import { UserProfileComponent } from './user-profile.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { EoNgCheckboxModule } from 'eo-ng-checkbox';
import { EoNgDropdownModule } from 'eo-ng-dropdown';
import { EoNgFeedbackMessageModule, EoNgFeedbackMessageService } from 'eo-ng-feedback';
import { EoNgInputModule } from 'eo-ng-input';
import { EoNgSelectModule } from 'eo-ng-select';
import { NzFormPatchModule } from 'ng-zorro-antd/core/form';
import { NzHighlightModule } from 'ng-zorro-antd/core/highlight';
import { NzNoAnimationModule } from 'ng-zorro-antd/core/no-animation';
import { NzOutletModule } from 'ng-zorro-antd/core/outlet';
import { NzOverlayModule } from 'ng-zorro-antd/core/overlay';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzGridModule } from 'ng-zorro-antd/grid';
import {  OverlayModule } from '@angular/cdk/overlay';
import { ScrollingModule } from '@angular/cdk/scrolling'



@NgModule({
  declarations: [
    ResetPswComponent,
    UserProfileComponent],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    NzFormModule,
    NzGridModule,
    EoNgInputModule,
    EoNgDropdownModule,
    HttpClientModule,
    ScrollingModule,
    OverlayModule,
    EoNgFeedbackMessageModule,
    EoNgSelectModule,
    EoNgCheckboxModule,
    NzOverlayModule,
    NzHighlightModule,
    NzNoAnimationModule,
    NzFormPatchModule,
    NzOutletModule,
  ],
  exports:[ResetPswComponent,
    UserProfileComponent
  ],
  providers:[EoNgFeedbackMessageService,
  ]
})
export class UserProfileModule { }
