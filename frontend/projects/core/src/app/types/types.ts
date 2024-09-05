/* eslint-disable camelcase */
export interface UserData{
    sex?:number
    avatar?:string
    email:string
    phone:string
    userName:string
    nickName:string
    roleIds:Array<string>
    noticeUserId:string
}

export interface UserListData extends UserData{
    id:number
    status:number
    lastLogin:string
    createTime:string
    updateTime:string
    operateDisable:boolean
    operator:string
}
