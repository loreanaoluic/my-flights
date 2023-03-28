import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailedInformationModalComponent } from './detailed-information-modal.component';

describe('DetailedInformationModalComponent', () => {
  let component: DetailedInformationModalComponent;
  let fixture: ComponentFixture<DetailedInformationModalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DetailedInformationModalComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DetailedInformationModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
