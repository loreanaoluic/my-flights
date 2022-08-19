import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpdateFlightModalComponent } from './update-flight-modal.component';

describe('UpdateFlightModalComponent', () => {
  let component: UpdateFlightModalComponent;
  let fixture: ComponentFixture<UpdateFlightModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UpdateFlightModalComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UpdateFlightModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
