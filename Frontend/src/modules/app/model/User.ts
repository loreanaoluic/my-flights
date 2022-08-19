export interface User{
    ID: number;
    Username : string;
	Password : string;
	EmailAddress : string;
	FirstName : string;
	LastName : string;
	Role : string;
	Banned : boolean;
	Deactivated : boolean;
	Reports : number;
	Points : number;
}