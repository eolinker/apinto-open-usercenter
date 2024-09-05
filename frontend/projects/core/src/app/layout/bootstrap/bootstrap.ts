import { UserAvatarComponent } from '../../component/user-avatar/user-avatar.component';
import { UserAvatarModule } from '../../component/user-avatar/user-avatar.module';
import { Injector } from '@angular/core';

// 当本项目作为插件导入apinto-dashboard时，本模块里的bootstrap方法作为立即执行插件的立即执行函数被执行
// 采用module的方式，user项目被引入时，每个模块创建的实例只有一个，模块之间的通信通过Angular服务实现，保持了Angular的一致性和可维护性
export async function bootstrap( props: any): Promise<void> {

  // 使用这些服务执行你需要的操作
  const {
    pluginSlotHub,
    pluginProvider,
    platformProvider
  } = props

  // 需要修改基座路由，使未定义的路由跳转至登录页->主页
  pluginProvider.setRouterConfig(true, {
    path: '',
    pathMatch: 'full',
    redirectTo: 'login'
  })


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