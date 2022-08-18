import { Component } from '@angular/core';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { AdminService } from 'src/modules/admin/services/admin.service';
import { ToastrService } from 'ngx-toastr';
import { Airline } from 'src/modules/app/model/Airline';

@Component({
  selector: 'app-update-airline-modal',
  templateUrl: './update-airline-modal.component.html',
  styleUrls: ['./update-airline-modal.component.scss']
})
export class UpdateAirlineModalComponent {
  airline: Airline;

  constructor(
    private adminService: AdminService,
    public modalRef: MdbModalRef<UpdateAirlineModalComponent>,
    private toastrService: ToastrService
  ) { }

  update() {
    if ((<HTMLInputElement>document.getElementById("name")).value == "") {

      this.toastrService.error("Please fill in all fields!");

    } else {
    
      this.airline.Name = (<HTMLInputElement>document.getElementById("name")).value;
      this.adminService.updateAirline(this.airline);
    }
  }
}
