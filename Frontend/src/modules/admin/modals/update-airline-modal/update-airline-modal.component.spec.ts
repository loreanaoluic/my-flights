import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpdateAirlineModalComponent } from './update-airline-modal.component';

describe('UpdateAirlineModalComponent', () => {
  let component: UpdateAirlineModalComponent;
  let fixture: ComponentFixture<UpdateAirlineModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UpdateAirlineModalComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UpdateAirlineModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
