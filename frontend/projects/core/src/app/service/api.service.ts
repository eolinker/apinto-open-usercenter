import { HttpClient, HttpErrorResponse, HttpHeaders, HttpParams } from '@angular/common/http'
import {  Injectable } from '@angular/core'
import { EoNgFeedbackMessageService } from 'eo-ng-feedback'
import { catchError, Observable, throwError } from 'rxjs'
import { environment } from '../../environments/environment'

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor (private message: EoNgFeedbackMessageService,
              private http:HttpClient,
  ) { }

  urlPrefix:string = environment.urlPrefix

  // 登录接口
  login (body?: any, params?: {[key:string]:any}) {
    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }
    if (params && params['query']) {
      params['query'] = JSON.stringify(params['query'])
    }
    const p = new HttpParams({
      fromObject: params
    })
    return this.http.post(this.urlPrefix + 'sso/login', body, {
      params: p, withCredentials: true
    }).pipe(catchError(this.handleError))
  }

  // 检查cookie是否合法,不合法则需要登录
  checkAuth (body?: any, params?: {[key:string]:any}) {
    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }
    if (params && params['query']) {
      params['query'] = JSON.stringify(params['query'])
    }
    const p = new HttpParams({
      fromObject: params
    })
    return this.http.post(this.urlPrefix + 'sso/login/check', body, {
      params: p, withCredentials: true
    }).pipe(catchError(this.handleError))
  }

  // 退出登录
  logout (body?: any, params?: {[key:string]:any}) {
    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }
    if (params && params['query']) {
      params['query'] = JSON.stringify(params['query'])
    }
    const p = new HttpParams({
      fromObject: params
    })
    return this.http.post(this.urlPrefix + 'sso/logout', body, {
      params: p, withCredentials: true
    }).pipe(catchError(this.handleError))
  }

  // 商业授权中激活时需要上传文件
  authPostWithFile (url: string, body?: any, params?: {[key:string]:any}): Observable<any> {
    for (const index in params) {
      if (typeof params[index] === 'string') {
        params[index] = params[index].trim()
      }
    }

    for (const index in body) {
      if (typeof body[index] === 'string') {
        body[index] = body[index].trim()
      }
    }

    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }
    const headers = new HttpHeaders()
    return this.http.post(this.urlPrefix + '_system/' + url, body, {
      headers,
      params: params,
      withCredentials: true
    })
      .pipe(
        catchError(this.handleError)
      )
  }

  // 商业授权相关的get接口
  authGet (url: string, params?: {[key:string]:any}): Observable<any> {
    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }
    if (params && params['query']) {
      params['query'] = JSON.stringify(params['query'])
    }

    for (const index in params) {
      if (typeof params[index] === 'string') {
        params[index] = params[index].trim()
      }
    }

    const p = new HttpParams({
      fromObject: params
    })
    return this.http.get(this.urlPrefix + '_system/' + url, {
      params: p,
      withCredentials: true
    })
      .pipe(
        // retry(3),

        catchError(this.handleError)
      )
  }

  get (url: string, params?: {[key:string]:any},options?:{apiPrefix?:boolean}): Observable<any> {
    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }
    if (params && params['query']) {
      params['query'] = JSON.stringify(params['query'])
    }
    params = this.underline(url, params)

    for (const index in params) {
      if (typeof params[index] === 'string') {
        params[index] = params[index].trim()
      }
    }

    const p = new HttpParams({
      fromObject: params
    })
    return this.http.get(this.urlPrefix +( options?.apiPrefix ?'api/': 'api/module/user/') + url, {
      params: p,
      withCredentials: true
    })
      .pipe(
        // retry(3),

        catchError(this.handleError)
      )
  }

  post (url: string, body?: any, params?: {[key:string]:any}): Observable<any> {
    for (const index in params) {
      if (typeof params[index] === 'string') {
        params[index] = params[index].trim()
      }
    }

    if (body && !(body instanceof FormData)) {
      for (const index in body) {
        if (typeof body[index] === 'string') {
          body[index] = body[index].trim()
        }
      }
    }

    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }

    body = !(body instanceof FormData) ? this.underline(url, body) : body
    params = this.underline(url, params)

    return this.http.post(this.urlPrefix + 'api/module/user/' + url, body, {
      params: params,
      withCredentials: true
    })
      .pipe(
        catchError(this.handleError)
      )
  }

  put (url:string, body?:any, params?: {[key:string]:any},options?:{apiPrefix?:boolean}): Observable<any> {
    for (const index in params) {
      if (typeof params[index] === 'string') {
        params[index] = params[index].trim()
      }
    }

    for (const index in body) {
      if (typeof body[index] === 'string') {
        body[index] = body[index].trim()
      }
    }
    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }

    body = this.underline(url, body)
    params = this.underline(url, params)

    return this.http.put(this.urlPrefix +( options?.apiPrefix ?'api/': 'api/module/user/') + url, body, {
      params: params,
      withCredentials: true
    })
      .pipe(
        catchError(this.handleError)
      )
  }

  delete (url:string, params?: {[key:string]:any}):Observable<any> {
    for (const index in params) {
      if (typeof params[index] === 'string') {
        params[index] = params[index].trim()
      }
    }

    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }

    params = this.underline(url, params)
    return this.http.delete(this.urlPrefix + 'api/module/user/' + url, { params: params })
      .pipe(
        catchError(this.handleError)
      )
  }

  patch (url:string, body?:any, params?: {[key:string]:any}): Observable<any> {
    for (const index in params) {
      if (typeof params[index] === 'string') {
        params[index] = params[index].trim()
      }
    }

    for (const index in body) {
      if (typeof body[index] === 'string') {
        body[index] = body[index].trim()
      }
    }
    if (params) { params['namespace'] = 'default' } else { params = { namespace: 'default' } }

    body = this.underline(url, body)
    params = this.underline(url, params)

    return this.http.patch(this.urlPrefix + 'api/module/user/' + url, body, {
      params: params,
      withCredentials: true
    })
      .pipe(
        catchError(this.handleError)
      )
  }

  handleError = (error: HttpErrorResponse) => {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error)
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error)
    }
    if (error.error.msg) {
      this.message.error(error.error.msg)
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Something bad happened; please try again later.'))
  }

  // 下划线转驼峰
  camel (data:any):any {
    if (typeof data !== 'object' || !data) return data
    if (Array.isArray(data)) {
      return (data as Array<any>).map((item:any) => { return this.camel(item) })
    }
    const newData:any = {}
    for (const key in data) {
      const newKey = key.replace(/_([a-z])/g, (p, m) => m.toUpperCase())
      newData[newKey] = this.camel(data[key])
    }
    return newData
  }

  // 驼峰转下划线,其中监控的status_4xx和status_5xx需要特殊处理
  underline (url:string, data:any) :any {
    if (url.startsWith('dynamic')) {
      return data
    }
    if (typeof data !== 'object' || !data) return data
    if (Array.isArray(data)) {
      return data.map(item => this.underline(url, item))
    }
    const newData:any = {}
    for (const key in data) {
      // 首字母不参与转换
      let newKey = key[0] + key.substring(1).replace(/([A-Z])/g, (p, m) => `_${m.toLowerCase()}`
      )
      newKey = key === 'status4xx' ? 'status_4xx' : (key === 'status5xx' ? 'status_5xx' : newKey)
      newData[newKey] = this.underline(url, data[key])
    }
    return newData
  }
}
