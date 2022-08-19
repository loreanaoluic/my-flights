import { Component } from '@angular/core';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { ToastrService } from 'ngx-toastr';
import { UserService } from 'src/modules/user/services/user.service';
import { User } from '../../model/User';

@Component({
  selector: 'app-edit-user-modal',
  templateUrl: './edit-user-modal.component.html',
  styleUrls: ['./edit-user-modal.component.scss']
})
export class EditUserModalComponent {
  user: User;

  constructor(
    private userService: UserService,
    public modalRef: MdbModalRef<EditUserModalComponent>,
    private toastrService: ToastrService
  ) { }

  update() {
    if ((<HTMLInputElement>document.getElementById("firstName")).value == ""
    || (<HTMLInputElement>document.getElementById("lastName")).value == ""
    || (<HTMLInputElement>document.getElementById("emailAddress")).value == "") {

      this.toastrService.error("Please fill in all fields!");

    } else {
    
      this.user.FirstName = (<HTMLInputElement>document.getElementById("firstName")).value;
      this.user.LastName = (<HTMLInputElement>document.getElementById("lastName")).value;
      this.user.EmailAddress = (<HTMLInputElement>document.getElementById("emailAddress")).value;
      this.userService.updateUser(this.user);
    }
  }

}
