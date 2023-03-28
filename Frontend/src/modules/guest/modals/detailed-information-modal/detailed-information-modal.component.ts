import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Flight } from 'src/modules/app/model/Flight';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';

@Component({
  selector: 'app-detailed-information-modal',
  templateUrl: './detailed-information-modal.component.html',
  styleUrls: ['./detailed-information-modal.component.scss']
})
export class DetailedInformationModalComponent {
  flight: Flight[];

  constructor(
    public modalRef: MdbModalRef<DetailedInformationModalComponent>,
    private router: Router,
  ) { }

}
