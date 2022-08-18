import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllFlightsPageComponent } from './all-flights-page.component';

describe('AllFlightsPageComponent', () => {
  let component: AllFlightsPageComponent;
  let fixture: ComponentFixture<AllFlightsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllFlightsPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllFlightsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
