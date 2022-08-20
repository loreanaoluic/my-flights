import { Component, OnInit } from '@angular/core';
import { Airline } from 'src/modules/app/model/Airline';
import { AdminService } from 'src/modules/admin/services/admin.service';
import { MdbModalRef, MdbModalService } from 'mdb-angular-ui-kit/modal';
import { NewAirlineModalComponent } from '../../modals/new-airline-modal/new-airline-modal.component';
import { UpdateAirlineModalComponent } from '../../modals/update-airline-modal/update-airline-modal.component';
import { JwtHelperService } from "@auth0/angular-jwt";

@Component({
  selector: 'app-all-airlines',
  templateUrl: './all-airlines.component.html',
  styleUrls: ['./all-airlines.component.scss'],
  providers: [MdbModalService]
})
export class AllAirlinesComponent implements OnInit {
  modalRef: MdbModalRef<NewAirlineModalComponent>
  airlines: Airline[] = [];
  term: string;
  currentRole : any

  constructor(
    private modalService: MdbModalService,
    private adminService: AdminService
  ) { }

  ngOnInit(): void {
    const tokenString = localStorage.getItem('userToken');
    if (tokenString) {
      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(tokenString);
      this.currentRole = info.role;
    }
    this.adminService.getAllAirlines().subscribe((response) => {
      this.airlines = response;
    });
  }

  openNewAirlineModal() {
    this.modalRef = this.modalService.open(NewAirlineModalComponent);
  }

  deleteAirline(id: number) {
    this.adminService.deleteAirline(id);
    window.location.reload();
  }

  openEditAirlineModal(airline: Airline) {
    this.modalRef = this.modalService.open(UpdateAirlineModalComponent, { data: { airline: airline }
    });
  }

  openAirlineReviewsModal(airline: Airline) {

  }

}
