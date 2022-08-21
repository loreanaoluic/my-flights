import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AirlineReviewPageComponent } from './airline-review-page.component';

describe('AirlineReviewPageComponent', () => {
  let component: AirlineReviewPageComponent;
  let fixture: ComponentFixture<AirlineReviewPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AirlineReviewPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AirlineReviewPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
