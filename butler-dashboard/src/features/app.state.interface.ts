export interface IUserState {
    token: string;
    firtName: string;
    lastName: string;
}

export default interface IAppState {
    user: IUserState
}
