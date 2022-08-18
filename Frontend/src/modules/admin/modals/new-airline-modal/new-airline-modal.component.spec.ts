import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewAirlineModalComponent } from './new-airline-modal.component';

describe('NewAirlineModalComponent', () => {
  let component: NewAirlineModalComponent;
  let fixture: ComponentFixture<NewAirlineModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewAirlineModalComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NewAirlineModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
