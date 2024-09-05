import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core'
import { AbstractControl, FormGroup, UntypedFormBuilder, ValidatorFn, Validators } from '@angular/forms'
import { EoNgFeedbackMessageService } from 'eo-ng-feedback'
import { NzSafeAny } from 'ng-zorro-antd/core/types'
import { defaultAutoTips } from '../../constant/conf'
import { UserData } from '../../types/types'
import { BaseInfoService } from '../../service/base-info.service'
import { ApiService } from '../../service/api.service'

@Component({
  selector: 'eo-ng-apinto-user-profile',
  templateUrl: './user-profile.component.html'
})
export class UserProfileComponent implements OnInit {
  @Input() userId:string = '' // 用户id
  @Input() type:string = '' // 操作类型 
  @Input() accessLink:string = ''
  @Input() nzDisabled:boolean = false // 权限
  @Output() eoCloseDrawer:EventEmitter<any> = new EventEmitter()
  validateForm:FormGroup = new FormGroup({})
  autoTips: Record<string, Record<string, string>> = defaultAutoTips

  rolesList:Array<any> = []
  editPage:boolean = false

  constructor (private message: EoNgFeedbackMessageService,
    private baseInfo:BaseInfoService,
    private fb: UntypedFormBuilder,
    private api:ApiService) {
    const { required, email } = EoNgMyValidators
    this.validateForm = this.fb.group({
      userName: ['', [required, Validators.pattern('^[a-zA-Z][a-zA-Z0-9/_]*')]],
      userNickName: ['', [required]],
      noticeUserId: [''],
      email: ['', [required, email]],
      role: [''],
      desc: ['']
    })
  }

  
  ngOnInit (): void {
    switch (this.type) {
      case 'editUser':
        this.getOtherUserProfile(this.userId)
        this.validateForm.controls['userName'].disable()
        if (this.userId === this.baseInfo.userId) {
          this.validateForm.controls['role'].disable()
        }
        this.getRolesList(false)
        break
      case 'addUser':
        this.getRolesList(false)
        break
      case 'editCurrentUser':
        this.getRolesList(true)
        this.getCurrentUserProfile()
        this.validateForm.controls['userName'].disable()
        this.validateForm.controls['role'].disable()
        break
    }
  }

  getCurrentUserProfile () {
    this.api.get('my/profile',{},{apiPrefix:true}).subscribe(async (resp: any) => {
      if (resp.code === 0) {
        this.validateForm.controls['userName'].setValue(resp.data.profile.userName)
        this.validateForm.controls['userNickName'].setValue(resp.data.profile.nickName)
        this.validateForm.controls['noticeUserId'].setValue(resp.data.profile.noticeUserId)
        this.validateForm.controls['email'].setValue(resp.data.profile.email)
        this.validateForm.controls['role'].setValue(resp.data.profile.roleIds[0])
        this.validateForm.controls['desc'].setValue(resp.data.describe)
      } else {
        this.message.error(resp.msg || '获取用户信息失败!')
      }
    })
  }

  getOtherUserProfile (id:string) {
      this.api.get('user/profile', { id: id }).subscribe((resp:{code:number, data:{profile:UserData, describe:string}, msg:string})=>{
        if (resp.code === 0) {
          this.validateForm.controls['userName'].setValue(resp.data.profile.userName)
          this.validateForm.controls['userNickName'].setValue(resp.data.profile.nickName)
          this.validateForm.controls['noticeUserId'].setValue(resp.data.profile.noticeUserId)
          this.validateForm.controls['email'].setValue(resp.data.profile.email)
          this.validateForm.controls['role'].setValue(resp.data.profile.roleIds[0])
          this.validateForm.controls['desc'].setValue(resp.data.describe)
        }
      })
  }

  // 获取角色id与title对应值, 传入list时,需要为该list的角色id与角色名匹配
  // 传入参数为true时,展示超管角色
  getRolesList (showM:boolean) {
    this.api.get('role/options').subscribe((resp:any)=>{
      if (resp.code === 0) {
        this.rolesList = showM
          ? resp.data.roles
          : resp.data.roles.filter((item:any) => {
            return item.title !== '超级管理员'
          })
        for (const index in this.rolesList) {
          this.rolesList[index].label = this.rolesList[index].title
          this.rolesList[index].value = this.rolesList[index].id
        }
        this.rolesList.push({ label: '未分配', value: '' })
      }
    })
  }


  backToList (value:any) {
    this.closeModal(value)
    this.eoCloseDrawer.emit(value)
  }

  // 当表单通过验证后,根据父组件data传来的type提交表单
  saveUserProfile () {
    if (this.validateForm.valid) {
      switch (this.type) {
        case 'addUser':
          this.api.post('user/profile',{
              userName: this.validateForm.value.userName,
              nickName: this.validateForm.value.userNickName,
              noticeUserId: this.validateForm.value.noticeUserId,
              email: this.validateForm.value.email,
              desc: this.validateForm.value.desc || '',
              roleIds: [this.validateForm.value.role]
                }).subscribe((resp:any)=>{
                  if (resp.code === 0) {
                    this.showMessageAndCloseModal(resp)
                  } else {
                    this.message.error(resp.msg || '修改失败!')
                  }
                })
            break
        case 'editUser':
          this.api.put('user/profile',{
              userName: this.validateForm.controls['userName'].value,
              nickName: this.validateForm.value.userNickName,
              noticeUserId: this.validateForm.value.noticeUserId,
              email: this.validateForm.value.email,
              desc: this.validateForm.value.desc || '',
              roleIds: [this.validateForm.value.role]
            }, {id:this.userId}).subscribe((resp:any)=>{
            if (resp.code === 0) {
              this.showMessageAndCloseModal(resp)
            } else {
              this.message.error(resp.msg || '修改失败!')
            }
                })
            break
        case 'editCurrentUser':
          this.api.put('my/profile', {
            nickName: this.validateForm.value.nickName,
            noticeUserId: this.validateForm.value.noticeUserId,
            email: this.validateForm.value.email,
            desc: this.validateForm.value.desc || ''
          },{},{apiPrefix:true}).subscribe((resp:any) => {
            if (resp.code === 0) {
              this.message.success(resp.msg || '修改成功!', { nzDuration: 1000 })
              this.closeModal()
            } else {
              this.message.error(resp.msg || '修改失败!')
            }
          })
          break;
      }
    } else {
      Object.values(this.validateForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty()
          control.updateValueAndValidity({ onlySelf: true })
        }
      })
    }
    return false
  }

  showMessageAndCloseModal(resp:any){
    this.message.success(resp.msg || '编辑用户信息成功!', { nzDuration: 1000 })
    this.closeModal(true)
  }
  closeModal:(value?:any)=>void = () => {

  }
}

// current locale is key of the MyErrorsOptions
export type EoNgMyErrorsOptions = { 'zh-cn': string; en: string } & Record<string, NzSafeAny>;
export type EoNgMyValidationErrors = Record<string, EoNgMyErrorsOptions>;

export class EoNgMyValidators extends Validators {
  static override minLength (minLength: number): ValidatorFn {
    return (control: AbstractControl): EoNgMyValidationErrors | null => {
      if (Validators.minLength(minLength)(control) === null) {
        return null
      }
      return { minlength: { 'zh-cn': `最小长度为 ${minLength}`, en: `MinLength is ${minLength}` } }
    }
  }

  static override maxLength (maxLength: number): ValidatorFn {
    return (control: AbstractControl): EoNgMyValidationErrors | null => {
      if (Validators.maxLength(maxLength)(control) === null) {
        return null
      }
      return { maxlength: { 'zh-cn': `最大长度为 ${maxLength}`, en: `MaxLength is ${maxLength}` } }
    }
  }

  static roleAccess (control:AbstractControl): EoNgMyValidationErrors | null {
    const value = control.value
    if (value.size > 0) {
      return null
    } else {
      return { roleAccess: { 'zh-cn': '角色权限不能为空', en: 'Not Empty' } }
    }
  }
}
