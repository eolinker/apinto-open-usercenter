import { NgModule ,} from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { EoNgAutoCompleteModule } from 'eo-ng-auto-complete';
import { EoNgButtonModule } from 'eo-ng-button';
import { EoNgCheckboxModule } from 'eo-ng-checkbox';
import { EoNgDropdownModule } from 'eo-ng-dropdown';
import { EoNgFeedbackAlertModule, EoNgFeedbackDrawerModule, EoNgFeedbackMessageModule, EoNgFeedbackMessageService, EoNgFeedbackTooltipModule } from 'eo-ng-feedback';
import { EoNgInputModule } from 'eo-ng-input';
import { EoNgSelectModule } from 'eo-ng-select';
import { EoNgSwitchModule } from 'eo-ng-switch';
import { EoNgTableModule } from 'eo-ng-table';
import { EoNgTreeModule } from 'eo-ng-tree';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { NzOutletModule } from 'ng-zorro-antd/core/outlet';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzResizableModule } from 'ng-zorro-antd/resizable';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzSpinModule } from 'ng-zorro-antd/spin';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { EoNgCopyModule } from 'eo-ng-copy';
import { EoNgLayoutModule } from 'eo-ng-layout'
import { EoNgMenuModule } from 'eo-ng-menu';
import { UserProfileModule } from '../component/user-profile/user-profile.module';

const importsModule = [
  CommonModule,
  RouterModule,
  EoNgTreeModule,
  EoNgInputModule,
  EoNgTableModule,
  EoNgButtonModule,
  EoNgFeedbackDrawerModule,
  EoNgFeedbackMessageModule,
  FormsModule,
  NzFormModule,
  ReactiveFormsModule,
  EoNgSelectModule,
  EoNgCheckboxModule,
  EoNgSwitchModule,
  EoNgAutoCompleteModule,
  NzCheckboxModule,
  NzResizableModule,
  NzTableModule,
  NzOutletModule,
  EoNgDropdownModule,
  NzPopconfirmModule,
  NzDividerModule,
  EoNgFeedbackTooltipModule,
  NzUploadModule,
  NzSpinModule,
  EoNgCopyModule,
  EoNgFeedbackAlertModule,
  EoNgLayoutModule,
  EoNgMenuModule,
  UserProfileModule
]

@NgModule({
  declarations: [
  ],
  imports: [
    ...importsModule
  ],
  exports: [
  ],
  providers:[
    EoNgFeedbackMessageService
  ]
})
export class LayoutModule { 
  constructor() {
  }}
