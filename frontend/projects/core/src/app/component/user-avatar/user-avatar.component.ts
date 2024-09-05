import { Component, OnInit } from '@angular/core'
import { EoNgFeedbackMessageService, EoNgFeedbackModalService } from 'eo-ng-feedback'
import { Router } from '@angular/router'
import { NzModalRef } from 'ng-zorro-antd/modal'
import { UserProfileComponent } from '../user-profile/user-profile.component'
import { ResetPswComponent } from '../reset-psw/reset-psw.component'
import { ApiService } from '../../service/api.service'
import { MODAL_SMALL_SIZE } from '../../constant/app.config'
import { BaseInfoService } from '../../service/base-info.service'

@Component({
  selector: 'eo-ng-apinto-user-avatar',
  templateUrl: './user-avatar.component.html',
  styleUrls: ['./user-avatar.component.scss']
})
export class UserAvatarComponent implements OnInit {
  userMenu: Array<any> = []
  nickName: string = ''
  userName: string = ''
  modalRef:NzModalRef | undefined
  appService:any = {
    setUserRoleId:(props:any)=>{},
    setUserId:(props:any)=>{}
  }
  constructor (private message: EoNgFeedbackMessageService,
                private modalService:EoNgFeedbackModalService,
                private router: Router,
                private apiService: ApiService,
                private baseInfoService:BaseInfoService
  ) {
  }

  ngOnInit (): void {
    this.userMenu = [
      {
        title: '用户设置',
        click: this.userSetting
      },
      {
        title: '修改密码',
        click: this.changeUserPsw
      },
      {
        title: '退出登录',
        click: this.logout
      }
    ]
    this.getCurrentUserProfile()
  }

  getCurrentUserProfile () {
    this.apiService.get('my/profile',{},{apiPrefix:true}).subscribe((resp:any) => {
      if (resp.code === 0) {
        this.nickName = resp.data.profile.nickName
        this.userName = resp.data.profile.userName
        resp.data.profile.roleIds?.length && (this.baseInfoService.userRoleId = resp.data.profile.roleIds[0])
        this.baseInfoService.userId = resp.data.profile.id
        this.baseInfoService.userProfile = resp.data.profile
      } else {
        this.message.error(resp.msg || '获取用户信息失败!')
      }
    })
  }

  userSetting = () => {
    this.openDrawer('editCurrentUser')
  }

  changeUserPsw = () => {
    this.openDrawer('changePsw')
  }

  openDrawer (usage:string) {
    switch (usage) {
      case 'editCurrentUser':
        this.modalRef = this.modalService.create({
          nzTitle: '用户设置',
          nzWidth: MODAL_SMALL_SIZE,
          nzContent: UserProfileComponent,
          nzComponentParams: {
            type: usage,
            closeModal: this.closeModal
          },
          nzOnOk: (component:UserProfileComponent) => {
            component.saveUserProfile()
            return false
          }
        })
        break
      case 'changePsw':
        this.modalRef = this.modalService.create({
          nzTitle: '修改密码',
          nzWidth: MODAL_SMALL_SIZE,
          nzContent: ResetPswComponent,
          nzComponentParams: { type: usage, userName: this.userName, closeModal: this.closeModal },
          nzOnOk: (component:ResetPswComponent) => {
            component.resetPsw()
            return false
          }
        })
        break
    }
  }

  logout = () => {
    this.apiService.logout().subscribe((resp:any) => {
      if (resp.code === 0) {
        this.router.navigate(['/', 'login'])
      } else {
        this.message.error(resp.msg || '退出登录失败!')
      }
    })
  }

  closeModal =() => {
    this.modalRef?.close()
  }
}
