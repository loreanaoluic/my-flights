export interface User{
    ID: number;
    Username : string;
	EmailAddress : string;
	FirstName : string;
	LastName : string;
	Role : string;
	Banned : boolean;
	Deactivated : boolean;
	Reports : number;
	Points : number;
	AccountBalance : number;
}