import { Component, OnInit } from '@angular/core';
import { AdminService } from '../../services/admin.service';
import { User } from 'src/modules/app/model/User';

@Component({
  selector: 'app-all-users',
  templateUrl: './all-users.component.html',
  styleUrls: ['./all-users.component.scss']
})
export class AllUsersComponent implements OnInit {
  users: User[] = [];
  term: string;

  constructor(
    private adminService: AdminService
  ) { }

  ngOnInit(): void {
    this.adminService.getAllUsers().subscribe((response) => {
      this.users = response;
    });
  }

  banUser(id: number) {
    this.adminService.banUser(id);
    window.location.reload();
  }

  unbanUser(id: number) {
    this.adminService.unbanUser(id);
    window.location.reload();
  }

}
