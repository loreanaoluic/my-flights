import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { NewReview } from 'src/modules/app/model/NewReview';
import { Review } from 'src/modules/app/model/Review';
import { AdminService } from '../../services/admin.service';
import { JwtHelperService } from "@auth0/angular-jwt";
import { ToastrService } from 'ngx-toastr';
import { UserService } from 'src/modules/user/services/user.service';

@Component({
  selector: 'app-airline-review',
  templateUrl: './airline-review.component.html',
  styleUrls: ['./airline-review.component.scss']
})
export class AirlineReviewComponent implements OnInit {
  airlineId: number;
  reviews: Review[] = [];
  currentUserId: number;
  currentUserUsername: string;
  averageRating: number = 0;
  hasTicket: false;

  constructor(
    private route: ActivatedRoute,
    private toastr: ToastrService,
    private adminService: AdminService,
    private userService: UserService
  ) { }

  ngOnInit(): void {

    this.route.queryParams
      .subscribe(params => {
        this.airlineId = params['id'];
      }
    );

    this.adminService.getReviewsByAirlineId(this.airlineId)
      .subscribe((response) => {
        this.reviews = JSON.parse(response);

        let allRatings = 0;
        this.reviews.forEach(function(review) {  
          allRatings += review.rating;
        });
    
        if (this.averageRating != 0) {
          let averageRating = allRatings / this.reviews.length;
          this.averageRating = Math.round((averageRating + Number.EPSILON) * 100) / 100
        }
    });

    const tokenString = localStorage.getItem('userToken');
    if (tokenString) {
      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(tokenString);

      this.currentUserId = info.Id;
      this.currentUserUsername = info.username;
    }
  }

  reportUser(userId: number) {
    this.adminService.reportComment(userId);
  }

  createReview() {
    let hasTicket = false;

    this.adminService.getAirlineById(this.airlineId).subscribe((response) => {
      let airlineName = response.Name;

      this.userService.getTicketsByUserId(this.currentUserId).subscribe((response) => {
        response.forEach(function(ticket) {  
          if (ticket.AirlineName == airlineName) {
            hasTicket = true;
          }
        });
        if (hasTicket == false) {
          this.toastr.error("You cannot leave a comment about a company whose services you have never used!")
        } else {
          let rating = 0;
          if((<HTMLInputElement>document.getElementById("1")).checked) {
            rating = 1;
          } else if((<HTMLInputElement>document.getElementById("2")).checked) {
            rating = 2;
          } else if((<HTMLInputElement>document.getElementById("3")).checked) {
            rating = 3;
          } else if((<HTMLInputElement>document.getElementById("4")).checked) {
            rating = 4;
          } else if((<HTMLInputElement>document.getElementById("5")).checked) {
            rating = 5;
          } else {
            this.toastr.error("Please enter rating!")
          }

          if ((<HTMLInputElement>document.getElementById("message")).value == "") {
            this.toastr.error("Please enter comment!")
          } else {
            const review: NewReview = {
              user_id: Number(this.currentUserId),
              username: this.currentUserUsername,
              message: (<HTMLInputElement>document.getElementById("message")).value,
              rating: Number(rating),
              airline_id: Number(this.airlineId)
            };

            this.adminService.createReview(review);
            window.location.reload();
          }
        }
      });
    });
  }

}
