import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ModuleFederationService {
  providerFromCore:any
  pluginEventHub: any
  pluginSlotHub: any
  store:any
  initialized:boolean = false
  constructor() {
  }
}
