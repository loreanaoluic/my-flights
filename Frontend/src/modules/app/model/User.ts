export interface User{
    Id: number;
    Username : string;
	EmailAddress : string;
	FirstName : string;
	LastName : string;
	Role : string;
	Banned : boolean;
	Deactivated : boolean;
	Reports : number;
	Points : number;
}