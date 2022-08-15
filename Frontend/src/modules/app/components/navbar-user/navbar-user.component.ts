import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-navbar-user',
  templateUrl: './navbar-user.component.html',
  styleUrls: ['./navbar-user.component.scss']
})
export class NavbarUserComponent implements OnInit {
  currentRole : any

  constructor(
    private authService : AuthService, 
    private router: Router
  ) { }

  ngOnInit(): void {
    this.currentRole = this.authService.getCurrentUser()?.dtype;
  }

  logout(){
    this.authService.logout();
    this.router.navigate(["login"]);
  }
}
