import { Router } from '@angular/router';
import { UserAvatarComponent } from '../../component/user-avatar/user-avatar.component';
import { UserAvatarModule } from '../../component/user-avatar/user-avatar.module';
import { Injector } from '@angular/core';

// 当本项目作为插件导入apinto-dashboard时，本模块里的bootstrap方法作为立即执行插件的立即执行函数被执行
// 采用module的方式，user项目被引入时，每个模块创建的实例只有一个，模块之间的通信通过Angular服务实现，保持了Angular的一致性和可维护性
export async function bootstrap( props: any): Promise<void> {

  // 使用这些服务执行你需要的操作
  const {
    pluginEventHub,
    pluginSlotHub,
    pluginProvider,
    platformProvider,
    closeModal,
    router,
  } = props

  // 需要修改基座路由，使未定义的路由跳转至登录页->主页
  pluginProvider.setRouterConfig(true, {
    path: '',
    pathMatch: 'full',
    redirectTo: 'login'
  })

  
  // 需要监听http响应，判断用户是否需要重新登录
  pluginEventHub.on('httpResponse', (eventData: any) => {
    const continueRes = checkAccess(eventData.res.body.code,  closeModal, router)
    return {
      data: eventData,
      continue: continueRes
    }
  })


  const checkAccess = (code: number,closeModal: Function, router: Router) => {
    switch (code) {
      case -3:
        setTimeout(() => {
          closeModal()
          if (!router.url.includes('/login')) {
            router.navigate(['/', 'login'], {queryParams: {callback: router.url}, queryParamsHandling: 'merge'})
          }
        }, 1000)
        return false
      default:
        return true
    }
  }

  // 渲染用户头像
  const ModuleRef = await platformProvider.getPlatformRef().bootstrapModule(UserAvatarModule, {ngZone: platformProvider.getNgZone()})
      .catch((error: any) => {
        console.error("Bootstrap error:", error);
        if (error instanceof Error) {
          console.error("Error message:", error.message);
          console.error("Error stack:", error.stack);
        } else {
          console.error("Error details:", error);
        }
        throw error;
      });
  const ModuleInject: Injector = ModuleRef.injector;
  pluginSlotHub.addSlot('renderAvatar', [UserAvatarComponent, {ModuleInject}])
}