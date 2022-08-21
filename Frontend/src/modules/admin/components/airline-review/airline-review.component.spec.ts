import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AirlineReviewComponent } from './airline-review.component';

describe('AirlineReviewComponent', () => {
  let component: AirlineReviewComponent;
  let fixture: ComponentFixture<AirlineReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AirlineReviewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AirlineReviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
