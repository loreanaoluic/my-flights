import { Component } from '@angular/core';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { AdminService } from 'src/modules/admin/services/admin.service';
import { ToastrService } from 'ngx-toastr';
import { NewAirline } from 'src/modules/app/model/NewAirline';

@Component({
  selector: 'app-new-airline-modal',
  templateUrl: './new-airline-modal.component.html',
  styleUrls: ['./new-airline-modal.component.scss']
})
export class NewAirlineModalComponent {

  constructor(
    private adminService: AdminService,
    public modalRef: MdbModalRef<NewAirlineModalComponent>,
    private toastrService: ToastrService
  ) { }

  createNew() {
    if ((<HTMLInputElement>document.getElementById("name")).value == "") {

      this.toastrService.error("Please fill in all fields!");

    } else {
    
      const airline: NewAirline = {
        Name: (<HTMLInputElement>document.getElementById("name")).value,
      };

      this.adminService.addNewAirline(airline);
      window.location.reload();

    }
  }
  
}
