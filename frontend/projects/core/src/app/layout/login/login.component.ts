import { Component, OnInit, ViewEncapsulation } from '@angular/core'
import { Router } from '@angular/router'
import { Subscription } from 'rxjs'
import { ApiService } from '../../service/api.service'
import { ModuleFederationService } from '../../service/module-federation.service'

@Component({
  selector: 'eo-ng-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class LoginComponent implements OnInit {
  isBusiness:boolean = true
  version:string = ''
  updateDate:string= ''
  powered:string =''
  private subscription: Subscription = new Subscription()
  constructor (
    private api: ApiService,
    private router: Router,
    private mfe:ModuleFederationService
  ) {}

  ngOnInit () {
    this.version = this.mfe.providerFromCore?.dashboardVersion.version
    this.updateDate = this.mfe.providerFromCore?.dashboardVersion.updateDate.split('-').join('')
    this.powered = this.mfe.providerFromCore?.dashboardVersion.powered
    this.api.checkAuth().subscribe((resp: any) => {
      if (resp.code === 0) {
        this.router.navigate([this.mfe.providerFromCore.mainPage()], { queryParamsHandling: 'merge' })
      }
    })
  }

  ngOnDestroy () {
    this.subscription.unsubscribe()
  }
}
