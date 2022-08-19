import { Component, OnInit } from '@angular/core';
import { JwtHelperService } from "@auth0/angular-jwt";
import { UserService } from 'src/modules/user/services/user.service';
import { User } from '../../model/User';
import { EditUserModalComponent } from '../../modals/edit-user-modal/edit-user-modal.component';
import { MdbModalRef, MdbModalService } from 'mdb-angular-ui-kit/modal';
import { Router } from '@angular/router';

@Component({
  selector: 'app-my-profile',
  templateUrl: './my-profile.component.html',
  styleUrls: ['./my-profile.component.scss'],
  providers: [MdbModalService]
})
export class MyProfileComponent implements OnInit {
  modalRef: MdbModalRef<EditUserModalComponent>
  user: User;

  constructor(
    private modalService: MdbModalService,
    private router: Router,
    private userService: UserService
  ) { }

  ngOnInit(): void {
    const tokenString = localStorage.getItem('userToken');
    if (tokenString) {
      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(tokenString);
      this.userService.getUserById(info.Id).subscribe((response) => {
        this.user = response;
        this.user.ID = info.Id
      });
    }
  }

  openEditUserModal() {
    this.modalRef = this.modalService.open(EditUserModalComponent, { data: { user: this.user }
    });
  }

  deactivateAccount() {
    this.userService.deactivateAccount(this.user.ID);
    this.router.navigate(["login"]);
  }

}
